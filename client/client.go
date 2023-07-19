package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lc-1010/OneBlogService/global"
	rmd "github.com/lc-1010/OneBlogService/internal/middleware/cmd_grpc"
	"github.com/lc-1010/OneBlogService/pkg/tracer"
	pb "github.com/lc-1010/OneBlogService/proto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	opt := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			rmd.UnaryContentTimeout(),
			rmd.ClientTracing(),
			rmd.UnaryInterceptor(),
		),
	}

	clientConn, err := GetClientConn(ctx, "localhost:8004", opt)
	if err != nil {
		log.Fatalf("GetClientConn err:%v", err)
	}
	defer clientConn.Close()
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1"+time.Now().Local().String(),
		"user-id", "some-test-user-id"+time.Now().Local().Format(time.Stamp),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequst{Name: "rust"})
	opts := []trace.SpanStartOption{}
	if err != nil {
		fmt.Println(err)
	} else {
		opts = []trace.SpanStartOption{
			trace.WithAttributes(
				attribute.String("request.name", "rust"),
				attribute.Int("response.count", len(resp.List)),
			),
		}
	}

	ctx, span := global.Tracer.Tracer("client1").Start(context.Background(), "GetTagList", opts...)
	defer span.End()
	log.Printf("resp:----%v,%v\n", resp, ctx)
	time.Sleep(10 * time.Second)
	//time.Sleep(time.Second * 20)
	defer func(ctx context.Context) {
		global.Tracer.ForceFlush(ctx)
		global.Tracer.Shutdown(ctx)
	}(ctx)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, opts...)
}

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}
}

func setupTracer() error {
	tracerPorvider, err := tracer.NewJaegerTrancer(
		"cligrpc",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = tracerPorvider
	return nil
}

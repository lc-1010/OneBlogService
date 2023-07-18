package main

import (
	"context"
	"log"

	"github.com/lc-1010/OneBlogService/cmd/gateway/internal/middleware"
	"github.com/lc-1010/OneBlogService/global"
	"github.com/lc-1010/OneBlogService/pkg/tracer"
	pb "github.com/lc-1010/OneBlogService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	opt := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			middleware.ClientTracing(),
			middleware.UnaryContentTimeout(),
		),
	}

	clientConn, _ := GetClientConn(ctx, "localhost:8004", opt)
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequst{Name: "rust"})
	log.Printf("resp:%v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, opts...)
}

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatal("init.setupTracer err:%v", err)
	}
}

func setupTracer() error {
	tracerPorvider, err := tracer.NewJaegerTrancer(
		"grpc",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = tracerPorvider
	return nil
}

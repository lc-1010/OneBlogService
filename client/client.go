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
	clientv3 "go.etcd.io/etcd/client/v3"
	resolver "go.etcd.io/etcd/client/v3/naming/resolver"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const SERVICE_NAME = "tag-service"

func main() {
	ctx := context.Background()

	opt := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(
			rmd.UnaryContentTimeout(),
			rmd.ClientTracing(),
			rmd.UnaryInterceptor(),
		),
	}

	clientConn, err := GetClientConn(ctx, SERVICE_NAME, opt)
	//clientConn, err := GetClientConnOld(ctx, "localhost:8004", opt)
	//fmt.Println("client ok", clientConn)
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
	req := &pb.GetTagListRequst{Name: "go"}

	resp, err := tagServiceClient.GetTagList(ctx, req)

	if err != nil {
		log.Printf("got err :%#v\n", err)
	}
	// if err != nil {
	// 	log.Printf("got err :%v\n", err)
	// } else {
	// 	resp = &pb.GetTagListReply{
	// 		List:  []*pb.Tag{},
	// 		Pager: &pb.Pager{},
	// 	}
	// }
	log.Printf("\nresp:----%v\n", resp)

	defer func(ctx context.Context) {
		global.Tracer.Shutdown(ctx)
	}(ctx)

	//RespSapn(ctx, resp)
}

func RespSapn(ctx context.Context, resp *pb.GetTagListReply) {

	opts := []trace.SpanStartOption{
		trace.WithAttributes(
			attribute.String("request.name", "rust"),
			attribute.String("resp.String", resp.String()),
		),
	}

	tr := global.Tracer.Tracer("client")
	_, span := tr.Start(ctx, "resp", opts...)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	fmt.Printf("\nmd--%#v\n", md)
	for k, v := range md {

		span.SetAttributes(attribute.StringSlice(k, v))
	}

	defer span.End()
	//fmt.Printf("ok RunSpan,%#v", ctx)
}

func GetClientConnOld(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("trage:->", target)
	return grpc.DialContext(ctx, target, opts...)
}
func GetClientConn(ctx context.Context, serviceName string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()

	config := clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 15,
	}
	cli, err := clientv3.New(config)
	//cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		return nil, err
	}
	//ff(cli)

	//log.Printf("cli:%#v\n", cli)
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		return nil, err
	}
	log.Printf("etcdResolver:%#v\n", etcdResolver)
	//target := fmt.Sprintf("%s:%d", "tag-service", 8004)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver))

	//log.Printf("etcdResolver:%#v\n", target)
	g, err := grpc.DialContext(context.Background(), "etcd:///blogServer/grpc/tag-service", opts...)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("--->%#v", g)

	return g, nil

}

func ff(cli *clientv3.Client) {
	res, err := cli.Get(context.Background(), "/etcd://blogServer/grpc/tag-service/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range res.Kvs {
		fmt.Printf("---->Key: %s, Value: %s\n", kv.Key, kv.Value)
	}
	fmt.Printf("res-->%#v", res)
}

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}
}

func setupTracer() error {
	tracerPorvider, err := tracer.NewJaegerTrancer(
		"myblog",
		"127.0.0.1",
		"6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = tracerPorvider
	return nil
}

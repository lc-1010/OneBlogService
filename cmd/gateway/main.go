package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/proxy/grpcproxy"
	"go.uber.org/zap"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/lc-1010/OneBlogService/proto"
	"github.com/lc-1010/OneBlogService/server"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

//使用grpc-gateway 分流不同协议

var port string

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}

const SERVICE_NAME = "tag-service"

func grpcHandlderFunc(grpcServer *grpc.Server, otherHander http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHander.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer()),
	}
	s := grpc.NewServer(opts...)

	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux() //使用WithTransportCredentials和insecure.NewCredentials()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)
	return gwmux
}

func runHttpServer() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	return serverMux
}
func RunServer1(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)
	httpMux.Handle("/", gwmux)

	cc, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"http://localhost:2379",
		},
		DialTimeout: time.Second * 2,
		Context:     cc,
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	//res, _ := etcdClient.Get(context.Background(), "abc", clientv3.WithPrefix())

	defer etcdClient.Close()
	logger, _ := zap.NewDevelopment()
	target := fmt.Sprintf("/etcdv3://blogServer/grpc/%s", SERVICE_NAME)
	grpcproxy.Register(logger, etcdClient, target, ":"+port, 60)

	return http.ListenAndServe(":"+port, grpcHandlderFunc(grpcS, httpMux))
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err:%v", err)
	}

}
func RunServer(prot string) error {
	httpMux := runHttpServer()

	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", grpcHandlderFunc(grpcS, httpMux))
	return http.ListenAndServe(":"+prot, grpcHandlderFunc(grpcS, gatewayMux))
}
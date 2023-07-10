package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	pb "github.com/lc-1010/OneBlogService/proto"
	"github.com/lc-1010/OneBlogService/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var grpcPort string
var httpPort string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "8001", "grpc启动端口")
	flag.StringVar(&httpPort, "http_port", "8002", "http启动端口")
	flag.Parse()
}

func RunHttpServer(port string) error {
	serveMux := http.NewServeMux() //多路复用

	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})
	return http.ListenAndServe(":"+port, serveMux)
}

func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// grpcurl -plaintext  localhost:8001 proto.TagService.GetTagList
// curl 127.0.0.1:8002/ping
func main() {

	errs := make(chan error)
	go func() {
		err := RunHttpServer(httpPort)
		if err != nil {
			errs <- err
		}
	}()
	go func() {
		err := RunGrpcServer(grpcPort)
		if err != nil {
			errs <- err
		}
	}()

	err := <-errs
	if err != nil {
		log.Fatalf("Run Server err:%v", err)
	}
}

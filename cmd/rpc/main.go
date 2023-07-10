package main

import (
	"log"
	"net"

	pb "github.com/lc-1010/OneBlogService/proto"
	"github.com/lc-1010/OneBlogService/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func main() {
	port = "8004"
	s := grpc.NewServer()
	tag := server.NewTagServer()
	pb.RegisterTagServiceServer(s, tag)
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)

	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve err:%v", err)
	}
}

package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/lc-1010/OneBlogService/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/proxy/grpcproxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestMain(t *testing.T) {
	//cc, cancel := context.WithTimeout(context.Background(), time.Second*12)
	//defer cancel()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"http://localhost:2379",
		},
		DialTimeout: time.Second * 2,
		//Context:     cc,
	})
	fmt.Println("ok")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("o1k")
	res, err := etcdClient.Get(context.Background(), "abc", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Printf("--->%v\n", res)
	fmt.Println("o21k")
	defer etcdClient.Close()
}

func TestRegist(t *testing.T) {
	endpoint := "0.0.0.0:8004"
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	cc, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"http://localhost:2379",
		},
		DialTimeout: time.Second * 5,
		Context:     cc,
	})

	if err != nil {
		fmt.Println(err)

	}

	//res, _ := etcdClient.Get(context.Background(), "abc", clientv3.WithPrefix())

	defer etcdClient.Close()
	logger, _ := zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))
	target := "/etcd://blogServer/grpc/abcd"
	fmt.Println(target)
	grpcproxy.Register(logger, etcdClient, target, ":"+"8004", 60)

	resp, err := etcdClient.Get(cc, "a", clientv3.WithFirstRev()...)
	if err != nil {
		log.Fatalf("get:%#v\n", err)
	}

	for _, kv := range resp.Kvs {
		fmt.Printf("key=%s, value=%s\n", kv.Key, kv.Value)
	}
	time.Sleep(time.Second * 70)
}

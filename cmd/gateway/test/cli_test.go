package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
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

package main

import (
	"fmt"
	"net"
	"sync"
	"testing"
	"time"
)

func BenchmarkBroadcaster(b *testing.B) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		b.Fatal(err)
	}
	defer listener.Close()

	// 启动广播器
	go broadcaster()

	// 启动客户端
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			conn, err := net.Dial("tcp", listener.Addr().String())
			if err != nil {
				b.Fatal(err)
			}
			defer conn.Close()

			// 发送消息到服务器
			fmt.Fprintf(conn, "Hello, %d\n", i)

			// 读取服务器发回的消息
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				b.Fatal(err)
			}
			msg := string(buf[:n])

			// 检查消息是否正确
			expect := fmt.Sprintf("Hello, %d\n", i)
			if msg != expect {
				b.Errorf("expect %q, but got %q", expect, msg)
			}
			conn.Close()
		}()
	}

	// 等待所有客户端完成测试
	wg.Wait()
	listener.Close()
}

func TestBroadcasterWithBlocking(t *testing.T) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	// 启动广播器
	go broadcaster()

	// 启动客户端
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", listener.Addr().String())
			if err != nil {
				t.Fatal(err)
			}
			defer conn.Close()

			// 模拟消息延迟和阻塞
			time.Sleep(5 * time.Second)

			// 发送消息到服务器
			fmt.Fprintf(conn, "Hello, %d\n", id)

			// 读取服务器发回的消息
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				t.Fatal(err)
			}
			msg := string(buf[:n])

			// 检查消息是否正确
			expect := fmt.Sprintf("Hello, %d\n", id)
			if msg != expect {
				t.Errorf("expect %q, but got %q", expect, msg)
			}
		}(i)
	}

	// 等待所有客户端完成测试
	wg.Wait()
}

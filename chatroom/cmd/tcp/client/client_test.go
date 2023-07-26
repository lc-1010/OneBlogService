package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"testing"
	"time"
)

func TestSendMsg(t *testing.T) {

	for i := 0; i < 100; i++ {

		go func(i int) {
			conn, err := net.Dial("tcp", ":2020")
			if err != nil {
				panic(err)
			}
			for j := 0; j < 1000; j++ {
				time.Sleep(time.Microsecond * 500)
				s := strings.NewReader(fmt.Sprintf("%s%d\n", "hello", i))
				if _, err := io.Copy(conn, s); err != nil {
					log.Fatal(err)
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
	log.Println("done")
}

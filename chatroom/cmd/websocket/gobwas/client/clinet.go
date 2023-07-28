package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, _, _, err := ws.Dial(ctx, "ws://localhost:8080/ws")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	err = wsutil.WriteClientText(conn, []byte("Hello WebSocket server"))
	if err != nil {
		log.Panicln(err)
	}

	c := make(chan os.Signal, 1)
	scanner := bufio.NewScanner(os.Stdin)

	data := make(chan string)
	hasClosed := map[string]bool{"hasClosed": false}
	go func() {
		defer close(data)
		for {
			msg, _, err := wsutil.ReadServerData(conn)
			//这里是for 等待所以会造成阻塞 所以无法主动退出
			if err != nil {
				if err == io.EOF {
					log.Println("server say bye")
					c <- os.Kill
					return
				} else {
					if !hasClosed["hasClosed"] {
						log.Printf("  err:%v", err)
					} else {
						return
					}
				}
			} else {
				data <- string(msg)
			}

		}

	}()

	go func() {
		for scanner.Scan() {
			input := scanner.Bytes()
			if string(input) == "quit" {
				log.Println("ok will quit wait...")
				c <- os.Kill
				return
			}
			if err := wsutil.WriteClientText(conn, input); err != nil {
				log.Printf("NewScanner %v\n", err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	signal.Notify(c, os.Interrupt)
t:
	for {
		select {
		case msg := <-data:
			fmt.Printf(">: %v\n", msg)
		case <-c:
			hasClosed["hasClosed"] = true
			break t
		}
	}

	log.Println("ok bye")

}

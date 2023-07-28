package main

import (
	"bufio"
	"context"
	"log"
	"os"
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

	err = wsutil.WriteClientText(conn, []byte("Hello WebSocket server"))
	if err != nil {
		log.Panicln(err)
	}
	go func() {
		for {
			msg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				log.Panicln(err)
			}
			log.Printf(">> %v\n", string(msg))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Bytes()
		if err := wsutil.WriteClientText(conn, input); err != nil {
			log.Printf("NewScanner %v\n", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
}

// func mustCopy(dst io.Writer, src io.Reader) {
// 	if _, err := io.Copy(dst, src); err != nil {
// 		log.Fatal(err)
// 	}
// }

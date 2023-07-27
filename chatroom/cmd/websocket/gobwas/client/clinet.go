package main

import (
	"context"
	"log"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, _, _, err := ws.Dial(ctx, "ws://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	err = wsutil.WriteClientText(conn, []byte("Hello WebSocket server"))
	if err != nil {
		log.Panicln(err)
	}
	msg, op, err := wsutil.ReadServerData(conn)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("recv: %v,%v", string(msg), op)
	conn.Close()
}

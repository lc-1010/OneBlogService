package main

import (
	"context"
	"log"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "server eroor")

	err = wsjson.Write(ctx, c, "Hello WebSocket server")
	if err != nil {
		panic(err)
	}
	var v any
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}
	log.Printf("recv: %v", v)
	c.Close(websocket.StatusNormalClosure, "")

}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "HTTP , Hello")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		conn, err := websocket.Accept(w, req, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "servr error")
		ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
		defer cancel()

		var v any
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("recv: %v", v)
		err = wsjson.Write(ctx, conn, "Hello WebSocket clinet")
		if err != nil {
			log.Println(err)
			return
		}

		conn.Close(websocket.StatusNormalClosure, "200")
	})
	log.Fatal(http.ListenAndServe(":2021", nil))
}

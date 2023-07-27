package main

import (
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			panic(err)
		}
		go func() {
			defer conn.Close()
			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("read:%v\n", err)
					break
				}
				log.Println(msg)

				err = wsutil.WriteServerMessage(conn, op, []byte("Hello   server"))
				if err != nil {
					log.Printf("write:%v\n", err)
					break
				}
			}
		}()
	}))
}

var (
	messageChannel    = make(chan string, 8)
	userIdListChannel = make(chan int, 1)
	closeChannel      = make(chan struct{})
	enteringChannel   = make(chan User)
	leavingChannel    = make(chan User)
)

type User struct {
	ID   int
	Name string
	Addr string
}

func broadcaster() {
	usersMap := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			userIdListChannel <- user.ID
		case user := <-leavingChannel:
			delete(usersMap, &user)
		}
	}
}

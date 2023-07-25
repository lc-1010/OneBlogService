package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var (
	enteringChannel   = make(chan *User)
	leavingChannel    = make(chan *User)
	messageChannel    = make(chan string, 8)
	userIdListChannel = make(chan int, 1)
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	go func() {
		userIdListChannel <- 0

	}()

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	usersMap := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			usersMap[user] = struct{}{}
		case user := <-leavingChannel:
			delete(usersMap, user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			for user := range usersMap {
				user.MessageChannel <- msg
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string),
	}
	go sendMessage(conn, user.MessageChannel)

	user.MessageChannel <- "Welcome!" + user.String()
	messageChannel <- "user: `" + strconv.Itoa(user.ID) + "` has enter"

	enteringChannel <- user

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- strconv.Itoa(user.ID) + ":" + input.Text()
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误:", err)
	}

	//
	leavingChannel <- user
	messageChannel <- "user: `" + strconv.Itoa(user.ID) + "` has leave"
}

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (user *User) String() string {
	return fmt.Sprintf("u-%d-%s", user.ID, strings.Split(user.Addr, ":")[0])
}

func GenUserID() int {
	last := <-userIdListChannel
	userIdListChannel <- last + 1
	return last + 1
}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

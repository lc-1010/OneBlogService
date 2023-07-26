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
	// 用户列表
	enteringChannel = make(chan *User)
	// 用户离开
	leavingChannel = make(chan *User)
	// messageChannel 缓冲区的大小是 8，这意味着 channel 可以存储最多 8 个字符串元素
	messageChannel = make(chan string, 2)
	// 增长id
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
	i := 0
	for {

		select {
		case user := <-enteringChannel:
			usersMap[user] = struct{}{}
		case user := <-leavingChannel:
			delete(usersMap, user)
			close(user.MessageChannel)
			//缓冲区的大小是 8，这意味着最多可以存储 8 条消息。
			//如果缓冲区已满，新的消息将会被阻塞，
			//直到有一个接收操作从 channel 中读取一个消息。
		case msg := <-messageChannel: //channel 空间限制住了
			i++
			msg = msg + "_" + fmt.Sprint(i)
			for user := range usersMap { //这里如果没有用户太多
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

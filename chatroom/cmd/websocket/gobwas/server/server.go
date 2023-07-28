package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

func main() {
	userIdListChannel <- 0
	go broadcaster()
	http.HandleFunc("/ws", handleConn)
	http.ListenAndServe(":8080", nil)
}

func getID() int {
	id := <-userIdListChannel
	userIdListChannel <- id + 1
	return id + 1
}

func handleConn(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		panic(err)
	}

	user := User{
		ID:   getID(),
		Name: getName(r.RemoteAddr),
		conn: &conn,
	}
	enteringChannel <- user
	wsutil.WriteServerMessage(conn, ws.OpText,
		[]byte(fmt.Sprintf("%s,%s!", "hello", user.Name)))

	go func(leavingChannel chan User) {
		defer func() {
			leavingChannel <- user
			conn.Close()
		}()
		for {
			msg, _, err := wsutil.ReadClientData(conn)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}
				break
			}
			log.Println(string(msg))
			newMsg := fmt.Sprintf("[%d-%s]:%s", user.ID, user.Name, string(msg))
			messageChannel <- []byte(newMsg)
		}

	}(leavingChannel)
}

var (
	messageChannel    = make(chan []byte, 8)
	userIdListChannel = make(chan int, 1)
	//closeChannel      = make(chan struct{})
	enteringChannel = make(chan User)
	leavingChannel  = make(chan User)
)

type User struct {
	ID   int
	Name string
	Addr string
	Msg  chan string
	conn *net.Conn
}

func broadcaster() {
	usersMap := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			//userIdListChannel <- user.ID 没有一值处理结果满了阻塞了
			usersMap[&user] = struct{}{}
			log.Println(user.ID, "enter", len(usersMap))
		case lu := <-leavingChannel:
			delete(usersMap, &lu)
			log.Println(lu.ID, "leave")
		case msg := <-messageChannel:
			log.Println("got messge", len(usersMap))
			for user := range usersMap {
				if strings.Split(strings.Trim(string(msg), "["), "-")[0] != strconv.Itoa(user.ID) {
					err := wsutil.WriteServerMessage(*user.conn, ws.OpText, []byte(msg))
					if err != nil {
						log.Printf("write:%v\n", err)
						delete(usersMap, user)
						continue
					}
				}
			}
			log.Println("send done")
		}

	}
}

func getName(s string) string {

	// 1. 提取数字
	numStr := strings.Split(s, ":")[3]
	num, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println(numStr, num, s)
	// 2. 转16进制
	hex := fmt.Sprintf("%X", num)

	// 3. 转整数
	val, _ := strconv.ParseInt(hex, 16, 32)
	// 4. 检查范围
	// 5. 映射到范围内
	valTmp := val>>10 + 0x1F600
	fmt.Println("tmp", valTmp)
	// 6. 检查返回的emoji是否合法
	for i := 0; i < 5; i++ {
		e, err := emoji.LookupEmoji(string(rune(val)))
		fmt.Printf("ok-->%#v %v\n", e, err)
		if err == nil {
			return e.Value
		} else {
			val = valTmp
		}

		val++

	}
	fmt.Println(val, "ooooo")
	// 8. 还是无法获取,返回默认emoji
	return string(rune(0x1F601))

}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ws(w http.ResponseWriter, req *http.Request) {

	conn, err := wu.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	c := &connection{conn, make(chan []byte, 256), &Data{}}

	h.r <- c

	// 向客户端发送消息
	go c.writer()
	// 接收客户端的消息 根据协议处理
	c.reader()

	// 最后用户断开连接 清理连接的用户信息
	defer func() {
		c.data.Type = "logout"
		userList = del(userList, c.data.User)
		c.data.UserList = userList
		dataJson, _ := json.Marshal(c.data)
		h.b <- dataJson
		h.r <- c
	}()
}

var userList []string

func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}

	defer c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.Type = "login"
			c.data.User = c.data.Content
			c.data.From = c.data.Content
			userList = append(userList, c.data.User)
			c.data.UserList = userList

			dataJson, _ := json.Marshal(c.data)
			h.b <- dataJson
		case "user":
			c.data.Type = "user"
			dataJson, _ := json.Marshal(c.data)
			h.b <- dataJson
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			c.data.UserList = userList

			dataJson, _ := json.Marshal(c.data)
			h.r <- c
			h.b <- dataJson
		default:
			fmt.Println("===========default==========")
		}
	}
}

func del(userListSlice []string, user string) []string {
	count := len(userListSlice)
	if count == 0 {
		return userListSlice
	}

	if count == 1 && userListSlice[0] == user {
		return []string{}
	}

	// 执行删除操作
	var copyUserList []string
	for i := range userListSlice {
		if userListSlice[i] == user {
			if i == count {
				return userListSlice[:i]
			}
			copyUserList = append(userList[:i], userList[i+1:]...)
		}
	}

	return copyUserList
}

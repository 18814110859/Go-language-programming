package webSocket

import (
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
	conn, err := wu.Upgrade(w, req, req.Header)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	c := &connection{conn, make(chan []byte, 256), &Data{}}

	// h.r <- c
}

func write() {
}

func read() {
}

func del() {
}

package main

import "encoding/json"

type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

var h = hub{
	make(map[*connection]bool),
	make(chan []byte),
	make(chan *connection),
	make(chan *connection),
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			// c.data.User = ""
			// c.data.Content = ""
			c.data.Type = "handshake"
			//userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataJson, _ := json.Marshal(c.data)
			c.sc <- dataJson
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}

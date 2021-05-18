package webSocket

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
		case data := <-h.b:
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.User = ""
			c.data.Type = ""
			c.data.Content = ""
			c.data.UserList = []string{}
			j, _ := json.Marshal(c.data)
			c.sc <- j
		case c := <-h.u:
			if _, ok := h.c[c]; ok {

			}
		}
	}
}

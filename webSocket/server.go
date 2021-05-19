package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", ws)
	err := http.ListenAndServe("127.0.0.1:3001", router)
	if err != nil {
		fmt.Println("err:", err)
	}
}

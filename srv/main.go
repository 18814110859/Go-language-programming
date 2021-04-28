package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

// 关于 & 和 * 的使用
type Server struct {
	Addr		string
	Port		int
	Conf 		*Config
}


type Config struct {
	Protocol	string
	Timeout		time.Duration
	MaxCount	int
	TLS			*tls.Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	return &Server{addr, port, conf}, nil
}


func main() {
	conf := Config{
		"TCP",
		30 * time.Second,
		1024,
		nil,
	}

	srv1, _ := NewServer("127.0.0.1", 9000, nil)
	srv2, _ := NewServer("127.0.0.1", 9000, &conf)
	fmt.Println(srv1)
	fmt.Println(srv2)
}




package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr		string
	Port		int
	Protocol	string
	Timeout		time.Duration
	MaxCount	int
	TLS			*tls.Config
}

type Option func(*Server)



func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxCount(maxCount int) Option {
	return func(s *Server) {
		s.MaxCount = maxCount
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}


func NewServer(addr string, port int, options ...func(server *Server)) (*Server, error) {
	// 设置默认值
	srv := Server{
		Addr:    	addr,
		Port:    	port,
		Protocol:	"tcp",
		Timeout:	30 * time.Second,
		MaxCount:	1024,
		TLS:		nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}


func main() {

	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(300 * time.Second), MaxCount(1024), TLS(nil))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}



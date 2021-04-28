package main

import (
	"crypto/tls"
	"errors"
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

// 使用一个builder 结构来做包装
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port

	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}


func (sb *ServerBuilder) WithMaxCount(maxCount int) *ServerBuilder {
	sb.Server.MaxCount = maxCount
	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() (Server, error) {
	return sb.Server, nil
}


func main() {
	sb := ServerBuilder{}
	server, err := sb.Create("127.0.0.1", 8080).
		WithProtocol("tcp").
		WithTimeout(60 * time.Second).
		WithMaxCount(1024).
		Build()

	if err != nil {
		errors.New("create server fill")
	}

	fmt.Printf("%p:%+v", &server, server)
}









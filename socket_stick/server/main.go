package server

import (
	"../proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// NewReader 返回一个新的 Reader，其缓冲区的大小为默认值。
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}

		fmt.Printf("收到client端发来的数据：%s", msg)

		// 发送数据
		conn.Write([]byte(msg))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listener.Close()

	for {
		// Accept等待并将下一个连接返回给侦听器。
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}

		go process(conn) // 启动一个 goroutine 处理连接
	}
}

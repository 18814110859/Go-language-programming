package client

import (
	"../proto"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp 连接失败, err:", err)
		return
	}
	// 关闭链接
	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("message encode field : err", err)
			return
		}
		conn.Write(data)
	}
}

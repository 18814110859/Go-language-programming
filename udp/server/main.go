package udpServer

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30001,
	})

	if err != nil {
		fmt.Println("listen failed , err:", err)
		return
	}

	defer listener.Close()

	for {
		data := make([]byte, 1024) // 新建一个1024 的slice 类型为 []byte

		// 接收数据
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}

		// 数据的数据
		fmt.Printf("接收的数据 data:%v addr:%v conn:%d\n", data[:n], addr, n)

		// 发送数据
		_, err1 := listener.Write(data[:n])
		if err1 != nil {
			fmt.Println("write udp failed, err:", err)
		}
	}

}

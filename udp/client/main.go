package udpClient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30001,
	})

	if err != nil {
		fmt.Println("连接失败, err: ", err)
		return
	}

	defer socket.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			continue
		}

		// 发送数据
		_, err = socket.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("发送数据失败, err:", err)
			continue
		}

		// 接收数据
		data := make([]byte, 4096)
		n, addr, err1 := socket.ReadFromUDP(data)
		if err1 != nil {
			fmt.Println("数据接收失败, err:", err1)
			continue
		}

		fmt.Printf("data:%v addr:%v count:%d", data[:n], addr, n)
	}

	/*
		// 发送数据
		sendData := []byte("hello world!")
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Println("发送数据失败，err:", err)
			return
		}
		// 接收数据
		data := make([]byte, 4096)
		n, addr, err1 := socket.ReadFromUDP(data)
		if err1 != nil {
			fmt.Println("数据接收失败, err:", err1)
			return
		}
		fmt.Printf("data:%v addr:%v count:%d", data[:n], addr, n)
	*/
}

package tcpClient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 一个TCP客户端进行TCP通信的流程如下：
// 1.建立与服务端的链接
// 2.进行数据收发
// 3.关闭链接
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp 连接失败, err:", err)
		return
	}
	// 关闭链接
	defer conn.Close()

	// 处理用户输入
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 数据处理
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 发送数据
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		// 接收数据
		buf := make([]byte, 512) // make 一个 512位 的 slice 类型为 []byte
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("reader failed, err:", err)
			return
		}

		fmt.Printf("接收的数据：%v", buf[:n])
	}
}

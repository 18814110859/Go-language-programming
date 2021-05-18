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

	// 测试粘包
	// 1.由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。
	// 简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
	// 2.接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。
	// 当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}

	return

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

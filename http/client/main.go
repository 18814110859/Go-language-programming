package client

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("127.0.0.1:8080")
	if err != nil {
		fmt.Println("request failed error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("request success Header:%v Status:%s", resp.Header, resp.Status)

	// 判断是否请求成功
	if resp.StatusCode == 200 {
		buf := make([]byte, resp.ContentLength)
		for {
			// 接收服务端信息
			n, err := resp.Body.Read(buf)
			if err != nil {
				fmt.Println("read failed err", err)
				return
			} else {
				fmt.Println(string(buf[:n]))
			}
		}
	}

}

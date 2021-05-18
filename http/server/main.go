package server

import (
	"fmt"
	"net/http"
)

func main() {

	// 127.0.0.1:8080/add
	http.HandleFunc("/add", addMethod)

	// 127.0.0.1:8080/remove
	http.HandleFunc("/remove", removeMethod)

	// 127.0.0.1:8080/get
	http.HandleFunc("/get", getMethod)

	// http.ListenAndServe
	// addr string：监听的地址
	// handler Handler：回调函数
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func getMethod(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s 连接成功 Method:%s Url:%s Header:%v Body:%v",
		request.RemoteAddr,
		request.Method,
		request.URL.Path,
		request.Header,
		request.Body)

	writer.Write([]byte("hello world go!"))
}

func addMethod(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("add success!"))
}

func removeMethod(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("remove success!"))
}

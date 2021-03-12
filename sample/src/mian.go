package sample

import (
	"log"
	_ "matchers"
	"os"
	"search"
)

// init 在main 之前调用
func init() {
	// 将日记输出到标准输出
	log.SetOutput(os.Stdout)
}

// main 是整个程序的入口
func main() {
	// 使用特定的项做搜索
	search.Run("president")
}

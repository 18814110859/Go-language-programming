package main

/**
为了能够构建这个工程，需要先把这个工程的根目录加入到环境变量GOPATH中。
假设calcproj 目录位于~/goyard下，则应编辑~/.bashrc文件，并添加下面这行代码：
export GOPATH=~/goyard/calcproj
source ~/.bashrc
GOPATH和PATH环境变量一样，也可以接受多个路径，并且路径和路径之间用冒号分割。
设置完GOPATH后，现在我们开始构建工程。
假设我们希望把生成的可执行文件放到 calcproj/bin目录中，需要执行的一系列指令如下：
*/

import "simplemath"
import (
	"os"
	"strconv"
)
import "fmt"

/**
使用说明
*/
var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquareroot of a non-negative value.")
}

func main() {
	args := os.Args
	// 如果输入的为空 或者 输入的小于两个则返回
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add": // 求加法
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		// 获取第一个 int的值
		v1, err1 := strconv.Atoi(args[2])
		// 获取第二个 int 的值
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt": // 求平方根
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer1>")
			return
		}
		// 获取第一个 int 的值
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer1>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}

package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 系统抛
func test01() {
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	//a[10] = 11
	index := 10
	a[index] = 10
	fmt.Println(a)
}

func test02() {
	getCircleArea(-5)
}

//
func test03() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

func test04() {
	test03()
	fmt.Println("test04")
}

func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		//panic("半径不能为负")
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

// 自定义error：
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf(`{"path":"%s","op":"%s","createTime":"%s","message":"%s"}`, p.path, p.op, p.createTime, p.message)
	// return fmt.Sprintf("{\"path\":\"%s\",\"op\":\"%s\",\"createTime\":\"%s\",\"message\":\"%s\"}\n", p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			createTime: fmt.Sprintf("%+v", time.Now()),
			message:    err.Error(),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()

	area, err := getCircleArea(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}

	err1 := Open("/Users/yu/code/Go-language-programming/error/test1.txt")

	switch v := err1.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}

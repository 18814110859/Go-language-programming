package main

import "fmt"

// Golang延迟调用：
// defer特性：
// 1. 关键字 defer 用于注册延迟调用。
// 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
// 3. 多个defer语句，按先进后出的方式执行。
// 4. defer语句中的变量，在defer声明时就决定了。
// defer用途：
// 1. 关闭文件句柄
// 2. 锁资源释放
// 3. 数据库连接释放

// go语言 defer
// go 语言的defer功能强大，对于资源管理非常方便，但是如果没用好，也会有陷阱。
// defer 是先进后出
// 这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Printf("%s closed\n", t.name)
}

func Close(t Test) {
	t.Close()
}

func test2 () {
	x, y := 10, 20
	defer func(i int) {
		println(i, y)
	} (x)

	x += 100
	y += 100
	println(x, y)
}


func main() {
	test2()
	

	//var whatever [5]struct{}
	//for i := range whatever {
	//	defer func() { fmt.Println(i) }()
	//}
	//ts := []Test{{"a"}, {"b"}, {"c"}}
	//for _, t := range ts {
	//	defer Close(t)
	//}



}

package main

import (
	"errors"
	"fmt"
)
import "sync"
import "time"

var lock sync.Mutex

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

func test2() {
	x, y := 10, 20
	defer func(i int) {
		println(i, y)
	}(x)

	x += 100
	y += 100
	println(x, y)
}

func testLock() {
	lock.Lock()
	lock.Unlock()
}

func testDeferLock() {
	lock.Lock()
	defer lock.Unlock()
}

// defer 与 closure
func foo(a, b int) (i int, err error) {

	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()

	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = 1 / 2
	return i, err
}

// 在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。
// 所以defer closure 输出结果为 2 而不是 1。
func foo1() (i int) {
	i = 1
	defer func() {
		fmt.Println(i)
	}()
	return 2
}

// defer nil 函数
func foo2() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func main() {
	test2()
	foo(100, 0)

	foo1()

	//var whatever [5]struct{}
	//for i := range whatever {
	//	defer func() { fmt.Println(i) }()
	//}
	//ts := []Test{{"a"}, {"b"}, {"c"}}
	//for _, t := range ts {
	//	defer Close(t)
	//}

	// 测试 defer 导致的性能问题
	func() {
		t1 := time.Now()
		for i := 0; i < 100000; i++ {
			testLock()
		}

		t2 := time.Since(t1)
		fmt.Println("testLock time: ", t2)
	}()

	func() {
		t1 := time.Now()
		for i := 0; i < 100000; i++ {
			testDeferLock()
		}
		t2 := time.Since(t1)
		fmt.Println("testDeferLock time", t2)
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	foo2()
}

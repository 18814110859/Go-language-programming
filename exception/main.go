package main

import (
	"errors"
	"fmt"
)

// 异常处理
// Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。

// 异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

// panic：
// 1、内置函数
// 2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
// 3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
// 4、直到goroutine整个退出，并报告错误

// recover：
// 1、内置函数
// 2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
// 3、一般的调用建议
// 		a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
// 		b). 可以获取通过panic传递的error

// 注意:
// 1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
// 2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
// 3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

// 由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
// func panic(v interface{})
// func recover() interface{}

// 除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。
// type error interface {
//	 Error() string
// }
// 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。

// 如何区别使用 panic 和 error 两种方式?
// 惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。

var ErrDivByZero = errors.New("division by zero")

func main() {
	// panicTest()
	// chanTest()
	// test()
	// test1()
	// test3()
	// test4(100, 10)
	// test5()

	// Go实现类似 try catch 的异常处理
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

func Try(fn func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fn()
}

func test5() {
	defer func() {
		fmt.Println(recover())
	}()

	switch z, err := div(10, 0); err {
	case nil:
		fmt.Println(z)
	case ErrDivByZero:
		panic(err)
	}
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}

	return x / y, nil
}

func panicTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("panic error!")
}

func chanTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan int, 10)
	close(ch)
	ch <- 1
}

func test4(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

func except() {
	fmt.Println(recover())
}

func test3() {
	defer except()
	//defer fmt.Println(recover())
	panic("test panic")
}

func test1() {

	defer func() {
		fmt.Println(recover())
	}()

	defer fmt.Println(recover())

	defer func() {
		func() {
			panic("defer inner")
		}()
	}()

	panic("test panic")
}

func test() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("last defer panic!")
	}()

	defer func() {
		panic("first defer panic!")
	}()

	panic("test panic!")
}

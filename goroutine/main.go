package main

/**
Go语言引入了goroutine概念，它使得并发编程变得非常简单。通过使用goroutine而不是裸用
操作系统的并发机制，以及使用消息传递来共享内存而不是使用共享内存来通信，Go语言让并
发编程变得更加轻盈和安全。
通过在函数调用前使用关键字go，我们即可让该函数以goroutine方式执行。goroutine是一种
比线程更加轻盈、更省资源的协程。Go语言通过系统的线程来多路派遣这些函数的执行，使得
每个用go关键字执行的函数可以运行成为一个单位协程。当一个协程阻塞的时候，调度器就会自
动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。而且调度的开销
非常小，一颗CPU调度的规模不下于每秒百万次，这使得我们能够创建大量的goroutine，从而可
以很轻松地编写高并发程序，达到我们想要的目的。
Go语言实现了CSP（通信顺序进程，Communicating Sequential Process）模型来作为goroutine
间的推荐通信方式。在CSP模型中，一个并发系统由若干并行运行的顺序进程组成，每个进程不
能对其他进程的变量赋值。进程之间只能通过一对通信原语实现协作。Go语言用channel（通道）
这个概念来轻巧地实现了CSP模型。channel的使用方式比较接近Unix系统中的管道（pipe）概念，
可以方便地进行跨goroutine的通信。
另外，由于一个进程内创建的所有goroutine运行在同一个内存地址空间中，因此如果不同的
goroutine不得不去访问共享的内存变量，访问前应该先获取相应的读写锁。Go语言标准库中的
sync包提供了完备的读写锁功能。
下面我们用一个简单的例子来演示goroutine和channel的使用方式。这是一个并行计算的例
子，由两个goroutine进行并行的累加计算，待这两个计算过程都完成后打印计算结果。
*/
import "fmt"

func main() {
	// 创建一个数组
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 创建一个切片
	resultChan := make(chan int, 2)
	// 开一个协程计算结果
	go sum(values[:len(values)/2], resultChan)
	// 开另一个协程计算结果
	go sum(values[len(values)/2:], resultChan)
	// 在 channel 中接收结果
	sum1, sum2 := <-resultChan, <-resultChan
	// 输出 sum1, sum2, sum1 + sum2
	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}

func sum(values []int, resultChan chan int) {
	// 初始化变量为0
	sum := 0
	// 计算结果
	for _, value := range values {
		sum += value
	}
	// 将结果发送到channel 中
	resultChan <- sum
}

package search

import (
	"log"
	"sync"
)

// TODO >> 创建一个 Matcher 的映射 （map）
// A map of registered matchers for searching
var matchers = make(map[string]Matcher)
//var matchers map[string]Matcher

/**
 * 执行搜索逻辑
 * 在 main 函数返回前，清理并终止所有之前启动的 goroutine
 * searchTerm string 搜索的搜索项
 */
func Run(searchTerm string) {

	// retrieve the list of feeds to search through
	// TODO >> 获取需要搜索的数据源列表
	// TODO >> 调用了 search 包的 RetrieveFeeds 函数，这个函数返回两个值。第一个返回值是一组 Feed 类型的切片。
	// TODO >> 切片是一种实现了一个动态数组的引用类型。在 Go 语言里可以用切片来操作一组数据。
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// TODO >> 创建一个无缓冲的通道，接收匹后的结果
	// TODO >> 声明并初始化该通道（channel）变量
	// TODO >> 通道本身实现的是一组带类型的值，这组值用于在 goroutine 之间传递数据。通道内置同步机制，从而保证通信安全。
	// create an unbuffered channel to receive match results to display
	results := make(chan *Result)

	// TODO >> 构建一个 waitGroup 以便处理所有的数据源
	// TODO >> 这个程序使用 sync 包的 WaitGroup 跟踪所有启动的 goroutine。
	// TODO >> 非常推荐使用 WaitGroup 来 跟踪 goroutine 的工作是否完成。
	// TODO >>WaitGroup 是一个计数信号量，我们可以利用它来统计所有的 goroutine 是不是都完成了工作。
	// setup await group so we can process all the feeds
	var waitGroup sync.WaitGroup


	// TODO >> 设置需要等待处理 每个数据源的 goroutines 的数量
	// set the number of go routines we need to wait for while
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// TODO >> 我们为每个数据源启动一个 goroutines 来查找结果（处理数据）
	// TODO >> 使用关键字 for range 对 feeds 切片做迭代
	// TODO >> 使用 for range 迭代切片时，每次迭代会返回两个值。第一个值是迭代的元素在切片里的索引位置，第二个值是元素值的一个副本
	// TODO >> 下划线标识符的作用是占位符，占据了保存 range 调用返回的索引值的变量的位置。如果要调用的函数返回多个值，而又不需要其中的某个值，就可以使用下划线标识符将其忽略。
	// launch a go routines for each feed to find the results
	for _, feed := range feeds {
		// TODO >> 获取一个匹配器用于查找
		// retrieve a matcher for the search
		// TODO >> 查找 map 里的键时，有两个选择：要么赋值给一个变量，要么为了精确查找，赋值给两个变量。赋值给两个变量时第一个值和赋值给一个变量时的值一样，是 map 查找的结果值。
		// TODO >> 如果指定了第二个值，就会返回一个布尔标志，来表示查找的键是否存在于 map 里。如果这个键不存在，map 会返回其值类型的零值作为返回值，如果这个键存在，map 会返回键所对应值的副本。
		matcher, exists := matchers[feed.Type]
		if !exists {
			// TODO >> 如果不存在，使用默认匹配器。这样程序在不知道对应数据源的具体类型时，也可以执行，而不会中断。
			matcher = matchers["default"]
		}

		// TODO >> 启动一个 goroutines 来执行搜索
		// TODO >> 关键字 go 启动了一个匿名函数作为 goroutine。匿名函数是指没有明确声明名字的函数。匿名函数也可以接受声明时指定的参数。
		// TODO >> 在 for range 循环里，我们为每个数据源，以 goroutine 的方式启动了一个匿名函数。这样可以并发地独立处理每个数据源的数据。
		// TODO >> 指针变量可以方便地在函数之间共享数据。使用指针变量可以让函数访问并修改一个变量的状态，而这个变量可以在其他函数甚至是其他 goroutine 的作用域里声明。
		// launch the go routines to perform the search
		go func(matcher Matcher, feed *Feed) {
			// TODO >> Match 函数的参数是一个 Matcher 类型的值、一个指向 Feed 类型值的指针、搜索项以及输出结果的通道。
			// TODO >> Match 函数会搜索数据源的数据，并将匹配结果输出到 results 通道。
			Match(matcher, feed, searchTerm, results)
			// TODO >> 每个goroutine完成其工作后，就会递减 WaitGroup 变量的计数值，当这个值递减到 0 时，我们就知道所有的工作都做完了。
			// TODO >> 一旦 Match 函数调用完毕，递减WaitGroup的计数。一旦每个goroutine 都执行调用 Match 函数和 Done 方法，程序就知道每个数据源都处理完成。
			// TODO >> 调用 Done方法这一行还有一个值得注意的细节：WaitGroup 的值没有作为参数传入匿名函数，但是匿名函数依旧访问到了这个值。
			// TODO >> 在匿名函数内访问 searchTerm 和 results变量，也是通过闭包的形式访问的。因为有了闭包，函数可以直接访问到那些没有作为参数传入的变量。
			// TODO >> 匿名函数并没有拿到这些变量的副本，而是直接访问外层函数作用域中声明的这些变量本身。
			// TODO >> 因为 matcher 和 feed 变量每次调用时值不相同，所以并没有使用闭包的方式访问这两个变量
			waitGroup.Done()
		} (matcher, feed)
	}

	// TODO >> 启动一个 goroutines 来监控是否所有的工作都做完了
	// TODO >> 需求随着每个 goroutine 搜索工作的运行，将结果发送到 results 通道，并递减 waitGroup 的计数，我们需要一种方法来显示所有的结果，并让 main 函数持续工作，直到完成所有的操作
	// TODO >> 我们以 goroutine 的方式启动了另一个匿名函数。这个匿名函数没有输入参数，使用闭包访问了 WaitGroup 和results 变量。
	// TODO >> 这个 goroutine 里面调用了 WaitGroup 的 Wait 方法。这个方法会导致 goroutine阻塞，直到 WaitGroup 内部的计数到达 0。
	// TODO >> 之后，goroutine 调用了内置的 close 函数，关闭了通道，最终导致程序终止。
	// launch a go routines to monitor when all the work is done
	go func () {
		// TODO >> 等待所有任务完成
		// wait for everything to be processed
		waitGroup.Wait()

		// TODO >> 关闭通道的方式，通知 Display 函数
		// TODO >> 退出程序
		// close the channel to signal to the Display
		// function that we exit the program
		close(results)
	} ()

	// TODO >> 这行调用了 match.go 文件里的 Display 函数。一旦这个函数返回，程序就会终止。
	// TODO >> 而之前的代码保证了所有 results 通道里的数据被处理之前，Display 函数不会返回。
	// start displaying results as they are available
	// return after the final result is displayed
	Display(results)
}

// register a matcher for use by the program
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
package main

// 单元测试
//不写测试的开发不是好程序员。我个人非常崇尚TDD（Test Driven Development）的，然而可惜的是国内的程序员都不太关注测试这一部分。 这篇文章主要介绍下在Go语言中如何做单元测试和基准测试。
//1.1. go test工具
//Go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。
//go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
//在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
//类型	格式	作用
//测试函数	函数名前缀为Test	测试程序的一些逻辑行为是否正确
//基准函数	函数名前缀为Benchmark	测试函数的性能
//示例函数	函数名前缀为Example	为文档提供示例文档
//go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
//Golang单元测试对文件名和方法名，参数都有很严格的要求。
//    1、文件名必须以xx_test.go命名
//    2、方法必须是Test[^a-z]开头
//    3、方法参数必须 t *testing.T
//    4、使用go test执行单元测试
//go test的参数解读：
//go test是go语言自带的测试工具，其中包含的是两类，单元测试和性能测试
//通过go help test可以看到go test的使用说明：
//格式形如： go test [-c] [-i] [build flags] [packages] [flags for test binary]

//参数解读：
//-c : 编译go test成为可执行的二进制文件，但是不运行测试。
//-i : 安装测试包依赖的package，但是不运行测试。
//关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空
//关于packages，调用go help packages，这些是关于包的管理，一般设置为空
//关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数
//-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。
//-test.run pattern: 只跑哪些单元测试用例
//-test.bench patten: 只跑那些性能测试用例
//-test.benchmem : 是否在性能测试的时候输出内存情况
//-test.benchtime t : 性能测试运行的时间，默认是1s
//-test.cpuprofile cpu.out : 是否输出cpu性能分析文件
//-test.memprofile mem.out : 是否输出内存性能分析文件
//-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件
//-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。
//你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。
//-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下
//-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。
//-test.timeout t : 如果测试用例运行时间超过t，则抛出panic
//-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理
//-test.short : 将那些运行时间较长的测试用例运行时间缩短

//目录结构：
//    test
//      |
//       —— calc.go
//      |
//       —— calc_test.go

// 测试函数
// 测试函数的格式
// 每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
// func TestName(t *testing.T){
//    // ...
// }

// 测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：\

// func TestAdd(t *testing.T){ ... }
// func TestSum(t *testing.T){ ... }
// func TestLog(t *testing.T){ ... }
// 其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：
// func (c *T) Error(args ...interface{})
// func (c *T) Errorf(format string, args ...interface{})
// func (c *T) Fail()
// func (c *T) FailNow()
// func (c *T) Failed() bool
// func (c *T) Fatal(args ...interface{})
// func (c *T) Fatalf(format string, args ...interface{})
// func (c *T) Log(args ...interface{})
// func (c *T) Logf(format string, args ...interface{})
// func (c *T) Name() string
// func (t *T) Parallel()
// func (t *T) Run(name string, f func(t *T)) bool
// func (c *T) Skip(args ...interface{})
// func (c *T) SkipNow()
// func (c *T) Skipf(format string, args ...interface{})
// func (c *T) Skipped() bool

// 测试组
// 我们现在还想要测试一下split函数对中文字符串的支持，这个时候我们可以再编写一个TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。

// 子测试
// 看起来都挺不错的，但是如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败了。我们可能会想到下面的解决办法

// 测试覆盖率
// 测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

// Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。例如：
//    split $ go test -cover
//    PASS
//    coverage: 100.0% of statements
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       0.005s
// 从上面的结果可以看到我们的测试用例覆盖了100%的代码。

// Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。例如：
//    split $ go test -cover -coverprofile=c.out
//    PASS
//    coverage: 100.0% of statements
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       0.005s
// 上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，
// 然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。

// 基准测试
// 基准测试函数格式
// 基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
//
// func BenchmarkName(b *testing.B){
//    // ...
// }

// 基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。
// testing.B拥有的方法如下：

// func (c *B) Error(args ...interface{})
// func (c *B) Errorf(format string, args ...interface{})
// func (c *B) Fail()
// func (c *B) FailNow()
// func (c *B) Failed() bool
// func (c *B) Fatal(args ...interface{})
// func (c *B) Fatalf(format string, args ...interface{})
// func (c *B) Log(args ...interface{})
// func (c *B) Logf(format string, args ...interface{})
// func (c *B) Name() string
// func (b *B) ReportAllocs()
// func (b *B) ResetTimer()
// func (b *B) Run(name string, f func(b *B)) bool
// func (b *B) RunParallel(body func(*PB))
// func (b *B) SetBytes(n int64)
// func (b *B) SetParallelism(p int)
// func (c *B) Skip(args ...interface{})
// func (c *B) SkipNow()
// func (c *B) Skipf(format string, args ...interface{})
// func (c *B) Skipped() bool
// func (b *B) StartTimer()
// func (b *B) StopTimer()

// 基准测试示例
// 基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试，输出结果如下：

// split $ go test -bench=Split
// goos: darwin
// goarch: amd64
// pkg: github.com/pprof/studygo/code_demo/test_demo/split
// BenchmarkSplit-8        10000000               203 ns/op
// PASS
// ok      github.com/pprof/studygo/code_demo/test_demo/split       2.255s
// 其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

// 我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。

// split $ go test -bench=Split -benchmem
// goos: darwin
// goarch: amd64
// pkg: github.com/pprof/studygo/code_demo/test_demo/split
// BenchmarkSplit-8        10000000               215 ns/op             112 B/op          3 allocs/op
// PASS
// ok      github.com/pprof/studygo/code_demo/test_demo/split       2.394s
// 其中，112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配。 我们将我们的Split函数优化如下：

// result = make([]string, 0, strings.Count(s, sep)+1)

// 这一次我们提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。我们来看一下这个改进会带来多大的性能提升：

//    split $ go test -bench=Split -benchmem
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/split
//    BenchmarkSplit-8        10000000               127 ns/op              48 B/op          1 allocs/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       1.423s

// 这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。

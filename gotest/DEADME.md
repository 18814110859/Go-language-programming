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


// 性能比较函数
// 上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，
// 比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？
// 再或者对于同一个任务究竟使用哪种算法性能最佳？
// 我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。

// 性能比较函数通常是一个带有参数的函数，被多个不同的 Benchmark 函数传入不同的值来调用。举个例子如下：

// func benchmark(b *testing.B, size int){/* ... */}
// func Benchmark10(b *testing.B){ benchmark(b, 10) }
// func Benchmark100(b *testing.B){ benchmark(b, 100) }
// func Benchmark1000(b *testing.B){ benchmark(b, 1000) }

// 运行基准测试：
//
//    split $ go test -bench=.
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//    BenchmarkFib1-8         1000000000               2.03 ns/op
//    BenchmarkFib2-8         300000000                5.39 ns/op
//    BenchmarkFib3-8         200000000                9.71 ns/op
//    BenchmarkFib10-8         5000000               325 ns/op
//    BenchmarkFib20-8           30000             42460 ns/op
//    BenchmarkFib40-8               2         638524980 ns/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/fib 12.944s
// 这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
//
// 最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。例如：
//
//    split $ go test -bench=Fib40 -benchtime=20s
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//    BenchmarkFib40-8              50         663205114 ns/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/fib 33.849s
// 这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了。


// 并行测试
// func (b B) RunParallel(body func(PB))会以并行的方式执行给定的基准测试。
//
// RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。
// 用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism。
// RunParallel通常会与-cpu标志一同使用。

// func BenchmarkSplitParallel(b *testing.B) {
//    // b.SetParallelism(1) // 设置使用的CPU数
//    b.RunParallel(func(pb *testing.PB) {
//        for pb.Next() {
//            Split("枯藤老树昏鸦", "老")
//        }
//    })
// }

// 执行一下基准测试：

// split $ go test -bench=.
// goos: darwin
// goarch: amd64
// pkg: github.com/pprof/studygo/code_demo/test_demo/split
// BenchmarkSplit-8                10000000               131 ns/op
// BenchmarkSplitParallel-8        50000000                36.1 ns/op
// PASS
// ok      github.com/pprof/studygo/code_demo/test_demo/split       3.308s
// 还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。

// Setup与TearDown
// 测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。


// TestMain
// 通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。

// 如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。
// TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。
//
// 一个使用TestMain来设置Setup和TearDown的示例如下：
//
// func TestMain(m *testing.M) {
//    fmt.Println("write setup code here...") // 测试之前的做一些设置
//    // 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//    retCode := m.Run()                         // 执行测试
//    fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
//    os.Exit(retCode)                           // 退出测试
// }
// 需要注意的是：在调用TestMain时, flag.Parse并没有被调用。
// 所以如果TestMain 依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。




// 子测试的Setup与Teardown
// 有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。



// Go怎么写测试用例
// 开发程序其中很重要的一点是测试，我们如何保证代码的质量，如何保证每个函数是可运行，运行结果是正确的，
// 又如何保证写出来的代码性能是好的，我们知道单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，
// 便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让线上的程序能够在高并发的情况下还能保持稳定。
// 本小节将带着这一连串的问题来讲解Go语言中如何来实现单元测试和性能测试。

// Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，
// testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例，
// 那么接下来让我们一一来看一下怎么写。

// 另外建议安装gotests插件自动生成测试代码:
//
// go get -u -v github.com/cweill/gotests/...

// 如何编写测试用例
// 由于go test命令只能在一个相应的目录下执行所有文件，所以我们接下来新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下。

// 接下来我们在该目录下面创建两个文件：gotest.go和gotest_test.go

// gotest.go:这个文件里面我们是创建了一个包，里面有一个函数实现了除法运算:

gotest_test.go:这是我们的单元测试文件，但是记住下面的这些原则：

文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码

你必须import testing这个包

所有的测试用例函数必须是Test开头

测试用例会按照源代码中写的顺序依次执行

测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态

测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。

函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。


```go
    package gotest
    import (
        "testing"
    )

    func Test_Division_1(t *testing.T) {
        if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
            t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
        } else {
            t.Log("第一个测试通过了") //记录一些你期望记录的信息
        }
    }

    func Test_Division_2(t *testing.T) {
        t.Error("就是不通过")
    }
```


我们在项目目录下面执行go test,就会显示如下信息：

```shell script
    --- FAIL: Test_Division_2 (0.00 seconds)
        gotest_test.go:16: 就是不通过
    FAIL
    exit status 1
    FAIL    gotest    0.013s
``` 



从这个结果显示测试没有通过，因为在第二个测试函数中我们写死了测试不通过的代码t.Error，那么我们的第一个函数执行的情况怎么样呢？默认情况下执行go test是不会显示测试通过的信息的，我们需要带上参数go test -v，这样就会显示如下信息：
```shell script
    === RUN Test_Division_1
    --- PASS: Test_Division_1 (0.00 seconds)
        gotest_test.go:11: 第一个测试通过了
    === RUN Test_Division_2
    --- FAIL: Test_Division_2 (0.00 seconds)
        gotest_test.go:16: 就是不通过
    FAIL
    exit status 1
    FAIL    gotest    0.012s
```

上面的输出详细的展示了这个测试的过程，我们看到测试函数1Test_Division_1测试通过，而测试函数2Test_Division_2测试失败了，最后得出结论测试不通过。接下来我们把测试函数2修改成如下代码：
```go
    func Test_Division_2(t *testing.T) {
        if _, e := Division(6, 0); e == nil { //try a unit test on function
            t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
        } else {
            t.Log("one test passed.", e) //记录一些你期望记录的信息
        }
    }
```
然后我们执行go test -v，就显示如下信息，测试通过了：
```shell script
    === RUN Test_Division_1
    --- PASS: Test_Division_1 (0.00 seconds)
        gotest_test.go:11: 第一个测试通过了
    === RUN Test_Division_2
    --- PASS: Test_Division_2 (0.00 seconds)
        gotest_test.go:20: one test passed. 除数不能为0
    PASS
    ok      gotest    0.013s
```

如何编写压力测试
压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：

压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
```
    func BenchmarkXXX(b *testing.B) { ... }
```
    
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数

在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行 文件名也必须以_test.go结尾

下面我们新建一个压力测试文件webbench_test.go，代码如下所示：

```go
package gotest

import (
    "testing"
)

func Benchmark_Division(b *testing.B) {
    for i := 0; i < b.N; i++ { //use b.N for looping 
        Division(4, 5)
    }
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
    b.StopTimer() //调用该函数停止压力测试的时间计数

    //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
    //这样这些时间不影响我们测试函数本身的性能

    b.StartTimer() //重新开始时间
    for i := 0; i < b.N; i++ {
        Division(4, 5)
    }
}
```

我们执行命令go test webbench_test.go -test.bench=".*"，可以看到如下结果：
```shell script
    Benchmark_Division-4                            500000000          7.76 ns/op         456 B/op          14 allocs/op
    Benchmark_TimeConsumingFunction-4            500000000          7.80 ns/op         224 B/op           4 allocs/op
    PASS
    ok      gotest    9.364s
```
上面的结果显示我们没有执行任何TestXXX的单元测试函数，显示的结果只执行了压力测试函数，第一条显示了Benchmark_Division执行了500000000次，每次的执行平均时间是7.76纳秒，第二条显示了Benchmark_TimeConsumingFunction执行了500000000，每次的平均执行时间是7.80纳秒。最后一条显示总共的执行时间。

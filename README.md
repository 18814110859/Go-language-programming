go语言的学习代码
go 的命令
    run  使用这个命令，会将编译，链接和运行3个步骤合并为一步
    build  只生成编译结果不自动运行
    version 查看版本信息
    
    

go 源代码安装

go 源代码升级



1.1.3. 优点
自带gc。

静态编译，编译好后，扔服务器直接运行。

简单的思想，没有继承，多态，类等。

丰富的库和详细的开发文档。

语法层支持并发，和拥有同步并发的channel类型，使并发开发变得非常方便。

简洁的语法，提高开发效率，同时提高代码的阅读性和可维护性。

超级简单的交叉编译，仅需更改环境变量。

1.1.4. Go语言的主要特征：
    1.自动立即回收。
    2.更丰富的内置类型。
    3.函数多返回值。
    4.错误处理。
    5.匿名函数和闭包。
    6.类型和接口。
    7.并发编程。
    8.反射。
    9.语言交互性。

1.1.6. Go语言命名：
- 1.Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：
    - 1）首字符可以是任意的Unicode字符或者下划线
    - 2）剩余字符可以是Unicode字符、下划线、数字
    - 3）字符长度不限

2.Go只有25个关键字
```
    break        default      func         interface    select
    case         defer        go           map          struct
    chan         else         goto         package      switch
    const        fallthrough  if           range        type
    continue     for          import       return       var
```

3.Go还有37个保留字

```
    Constants:    true  false  iota  nil
    Types:    int  int8  int16  int32  int64  
              uint  uint8  uint16  uint32  uint64  uintptr
              float32  float64  complex128  complex64
              bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover
```

1.1.7. Go语言声明：

- 有四种主要声明方式：
    - var（声明变量）, const（声明常量）, type（声明类型） ,func（声明函数）。

- Go的程序是保存在多个.go文件中，文件的第一行就是package XXX声明，
  用来说明该文件属于哪个包(package)，package声明下来就是import声明，
  再下来是类型，变量，常量，函数的声明。


1. Golang内置类型和函数
1.1. 内置类型
1.1.1. 值类型：
```
    bool
    int(32 or 64), int8, int16, int32, int64
    uint(32 or 64), uint8(byte), uint16, uint32, uint64
    float32, float64
    string
    complex64, complex128
    array    -- 固定长度的数组
```
1.1.2. 引用类型：(指针类型)
```
    slice   -- 序列数组(最常用)
    map     -- 映射
    chan    -- 管道
```

1.2. 内置函数
Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，或必须用于系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持。
```
    append          -- 用来追加元素到数组、slice中,返回修改后的数组、slice
    close           -- 主要用来关闭channel
    delete            -- 从map中删除key对应的value
    panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
    recover         -- 允许程序定义goroutine的panic动作
    real            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
    imag            -- 返回complex的虚部
    make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
    new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
    cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy            -- 用于复制和连接slice，返回复制的数目
    len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
    print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包
```
1.3. 内置接口error
```
    type error interface { //只要实现了Error()函数，返回值为String的都实现了err接口
        Error()    String
    }
```


1. Init函数和main函数
1.1. init函数
go语言中init函数用于包(package)的初始化，该函数是go语言的一个重要特性。

有下面的特征：

    1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等

    2 每个包可以拥有多个init函数

    3 包的每个源文件也可以拥有多个init函数

    4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)

    5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序

    6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
1.2. main函数
    Go语言程序的默认入口函数(主函数)：func main()
    函数体用｛｝一对括号包裹。

    func main(){
        //函数体
    }
1.3. init函数和main函数的异同
    相同点：
        两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用。
    不同点：
        init可以应用于任意包中，且可以重复定义多个。
        main函数只能用于main包中，且只能定义一个。
两个函数的执行顺序：

对同一个go文件的init()调用顺序是从上到下的。

对同一个package中不同文件是按文件名字符串比较“从小到大”顺序调用各文件中的init()函数。

对于不同的package，如果不相互依赖的话，按照main包中"先import的后调用"的顺序调用其包中的init()，如果package存在依赖，则先调用最早被依赖的package中的init()，最后调用main函数。

如果init函数中使用了println()或者print()你会发现在执行过程中这两个不会按照你想象中的顺序执行。这两个函数官方只推荐在测试环境中使用，对于正式环境不要使用。
    
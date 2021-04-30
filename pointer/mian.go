package main

import "fmt"

/**
 * Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。
 * 传递数据使用指针，而无须拷贝数据。类型指针不能进行偏移和运算。Go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）。
 * 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
 * Go语言中使用&字符放在变量前面对变量进行“取地址”操作。
 * Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等。
 */
func main() {
	a := 10
	c := &a
	fmt.Printf("a:%d %v ptr:%p c:%p\n", a, a, &a, c)
	modify(&a)
	fmt.Printf("a:%d %v ptr:%p c:%p\n", a, a, &a, c)

	// TODO>> 建立一个空指针
	// TODO>> 当一个指针被定义后没有分配到任何变量时，它的值为 nil
	var strPtr *string
	var intPtr *int
	fmt.Printf("strPtr:%v %p\n", strPtr, strPtr)
	fmt.Printf("intPtr:%v %p\n", intPtr, intPtr)

	// TODO>> 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	// TODO>> 指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。可以使用内置的new函数对 未分配内存空间的指针 进行初始化之后就可以正常对其赋值了
	strPtr = new(string)
	intPtr = new(int)
	*strPtr = "string" // 引发panic
	*intPtr = 100  // 引发panic

	// TODO >> 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
	// TODO >> 要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

	fmt.Printf("strPtr:%v %p\n", strPtr, strPtr)
	fmt.Printf("intPtr:%v %p\n", intPtr, intPtr)

	// 空指针的判断
	//if strPtr == nil {
	//	fmt.Printf("strPtr指针为nil, value:%v/n", strPtr)
	//}

	// TODO >> new是一个内置的函数，它的函数签名如下：
	// TODO >> func new(Type) *Type
	// TODO >> 1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
	// TODO >> 2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
	x := new(int)
	y := new(string)
	fmt.Printf("x:%v %d %T %p %p\n", x, *x, x, x, &x)
	fmt.Printf("y:%v %s %T %p %p\n", y, *y, y, y, &y)


	// TODO >> make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 为啥不能是array 用make 创建 因为array 是值不是引用

	// TODO >> 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
	// TODO >> make函数的函数签名如下：
	// TODO >> func make(t Type, size ...IntegerType) Type
	// TODO >> make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
	// var b map[string]int只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：
	var m map[string]string
	m = make(map[string]string)
	//m["key"] = "value"
	fmt.Printf("%v %T %p\n", m, m, &m)

	// TODO >> new与make的区别
	// TODO >> 1.二者都是用来做内存分配的。
	// TODO >> 2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	// TODO >> 3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。


	var g int
	var p *int
	p = &g
	fmt.Printf("g: %d %p %p %T %T %d\n", g, &g, p, g, p, *p)

	var str string
	var strP *string
	strP = &str
	fmt.Printf("str: %s %p %p %T %T %s\n", str, &str, strP, str, strP, *strP)

	var m1 map[string]string
	m1 = make(map[string]string)
	// m1 := make(map[string]string)

	var mapP *map[string]string
	mapP = &m1
	fmt.Printf("m1: %p %p %p %p\n", m1 , *mapP, &m1, mapP)
}


func modify(x *int) {
	*x = 100
}


package main

import "fmt"

// Golang for支持三种循环方式，包括类似 while 的语法。
// for循环是一个循环控制结构，可以执行指定次数的循环。

// 语法
// Go语言的For循环有3中形式，只有其中的一种使用分号。

// for init; condition; post { }
// for condition { }
// for { }
// init： 一般为赋值表达式，给控制变量赋初值；
// condition： 关系表达式或逻辑表达式，循环控制条件；
// post： 一般为赋值表达式，给控制变量增量或减量。

// for语句执行过程如下：
// ①先对表达式 init 赋初值；
// ②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；
// 否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

func main() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}


	// 循环语句range
	// Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	// for key, value := range oldMap {
	// 		newMap[key] = value
	// }

	d := "abc"
	// 忽略全部返回值，仅迭代。
	for range d {
	}

	a := [3]int{0, 1, 2}
	for i, v := range a { // index、value 都是从复制品中取出。
		if i == 0 { // 在修改前，我们先修改原数组。
			(&a)[1], (&a)[2] = 999, 999
		}
		(&a)[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
	}
	fmt.Println(a)


	s := []int{1, 2, 3, 4, 5}
	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}
	fmt.Printf("%v\n", s)

	// 两种引用类型 map、channel 是指针包装，而不像 slice 是 struct。

	// for 和 for range有什么区别?
	// 主要是使用场景不同
	// for可以 遍历array和slice 遍历key为整型递增的map 遍历string
	// for range可以完成所有for可以做的事情，却能做到for不能做的，包括 遍历key为string类型的map并同时获取key和value 遍历channel
	
	// 循环控制Goto、Break、Continue
	// 循环控制语句
	// 循环控制语句可以控制循环体内语句的执行过程。
	// GO 语言支持以下几种循环控制语句：
	// Goto、Break、Continue
	// 1.三个语句都可以配合标签(label)使用
	// 2.标签名区分大小写，定以后若不使用会造成编译错误
	// 3.continue、break配合标签(label)可用于多层循环跳出
	// 4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同

}

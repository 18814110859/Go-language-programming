package main

import (
	"fmt"
)

func main() {
	//var x Mover
	//var wangcai = &Dog{}
	//x = wangcai

	//wangcai.Move()
	//x.Move()

	//var fugui = &Dog{}
	//x = fugui
	//fugui.Move()
	//x.Move()

	//var s Sayer
	//var d = &Dog{}
	//s = d
	//s.Say()
	//d.Say()

	//var c = &Cat{}
	//s = c
	//s.Say()
	//c.Say()

	//var p People
	//var peo = &Student{}

	var m Mover
	var s Sayer
	var w WashingMachine

	// 一个类型实现多个接口 例子
	wangcai := &Dog{"旺财"}
	m = wangcai
	m.Move()
	s = wangcai
	s.Say()

	// 多个类型实现同一接口
	m = &Car{"宝马"}
	m.Move()

	m = &Cat{"小花"}
	m.Move()

	// 类型中嵌入其他类型或者结构体
	w = &Haier{"海尔洗衣机", Dryer{}}
	w.Dry()
	w.Wash()

	w = &Dryer{}
	w.Dry()
	w.Wash()

	// 接口嵌套
	var zoo Animal
	zoo = &Cat{name: "小花"}
	zoo.Say()
	zoo.Move()
	zoo = &Dog{name: "旺财"}
	zoo.Say()
	zoo.Move()

	var peo People = &Student{}
	fmt.Println(peo.Speak("bitch"))
	// 公式拆解
	// var peo People
	// var tem = &Student{}
	// peo = tem
	// peo.Speak("bitch")

	// 空接口 定义
	// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	// 空接口类型的变量可以存储任意类型的变量。

	// 类型断言
	// 接口值
	// 一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。
	// 这两部分分别称为接口的动态类型和动态值。
	// 想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
	//    x.(T)
	// 其中：
	// x：表示类型为interface{}的变量
	// T：表示断言x可能是的类型。
	// 该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。

}

// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数。
func show(a interface{}) {
	fmt.Printf("%+v\n", a)
}

// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。
func getStudentInfo() map[string]interface{} {
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "李白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	return studentInfo
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (s *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

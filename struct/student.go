package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与JSON序列化
// JSON(JavaScript Object Notation) 是一种轻量级的数据交换格式。
// 易于人阅读和编写。同时也易于机器解析和生成。
// JSON键值对是用来保存JS对象的一种方式，键/值对组合中的键名写在前面并用双引号""包裹，使用冒号:分隔，然后紧接着值； 多个键值之间使用英文,分隔 。

// 结构体标签（Tag）
// Tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来。
// Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
// 结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
// 注意事项： 为结构体编写Tag时，必须严格遵守键值对的规则。
// 结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
// 例如不要在key和value之间添加空格。
type Student struct {
	Id     int 		`json:"id"` // 通过指定tag实现json序列化该字段时的key
	Name   string	`json:"name"` // 通过指定tag实现json序列化该字段时的key
	Gender string	`json:"gender"`	// 通过指定tag实现json序列化该字段时的key
}

type Class struct {
	Title   string		`json:"title"`
	Student []*Student	`json:"student1"`
}

func main() {
	c := &Class{
		Title:   "101",
		Student: make([]*Student, 0, 10),
	}

	for i := 0; i < 10; i++ {
		stu := &Student{
			Id:     i,
			Name:   fmt.Sprintf("stu%2d", i),
			Gender: "man",
		}

		c.Student = append(c.Student, stu)
	}
	fmt.Printf("%#v\n", c)
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("JSON:%s\n", data)


	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Student":[{"Id":0,"Name":"stu 0","Gender":"man"},{"Id":1,"Name":"stu 1","Gender":"man"},{"Id":2,"Name":"stu 2","Gender":"man"},{"Id":3,"Name":"stu 3","Gender":"man"},{"Id":4,"Name":"stu 4","Gender":"man"},{"Id":5,"Name":"stu 5","Gender":"man"},{"Id":6,"Name":"stu 6","Gender":"man"},{"Id":7,"Name":"stu 7","Gender":"man"},{"Id":8,"Name":"stu 8","Gender":"man"},{"Id":9,"Name":"stu 9","Gender":"man"}]}`
	c1 := &Class{}
	// 此参数必须具有指针类型
	// 检查信息:分析对json.Unmarshal等函数的调用。如果传递来存储结果的参数不是指针或接口，则报告问题。这样的调用肯定会失败并返回错误。
	err1 := json.Unmarshal([]byte(str), c1)
	if err1 != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	//fmt.Printf("%#v\n", c1)
}

package main

import "fmt"

// Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
// 将MyInt定义为int类型
type MyInt int

// 类型别名是Go1.9版本添加的新功能。
// 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
// 我们之前见过的rune和byte就是类型别名，他们的定义如下：
type byte = uint8
type rune = int32

// 类型定义和类型别名的区别
// 结构体
// Go语言中的基础数据类型可以表示一些事物的基本属性，
// 但是当我们想表达一个事物的全部或部分属性时，
// 这时候再用单一的基本数据类型明显就无法满足需求了，
// Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
// 这种数据类型叫结构体，英文名称struct。
// 也就是我们可以通过struct来定义自己的类型了。

// 结构体的定义
// 使用type和struct关键字来定义结构体，具体代码格式如下：
type person struct {
	name string
	city string
	age  int8
}

// 1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 3.字段类型：表示结构体字段的具体类型。

type student struct {
	name string
	age  int
}

// 任意类型添加方法
// 在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
// 举个例子，我们基于内置的int类型使用 type 关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个MyInt。")
}

// 结构体的匿名字段
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
type person1 struct {
	string
	int
}

// 嵌套结构体（嵌套匿名结构体）
// 一个结构体中可以嵌套包含另一个结构体或结构体指针。
// Address 地址结构体
type Address struct {
	Province string
	City     string
}

// User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address // 匿名结构体
}

// 嵌套结构体的字段名冲突
// 嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。

// 结构体的“继承”
// Go语言中使用结构体也可以实现其他编程语言中面向对象的继承。

// 定义一个动物的结构属性
type Animal struct {
	name string
}
// Dog 狗 继承至动物的属性
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}
// 结构体字段的可见性
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

func main() {

	// 结构体实例化(匿名)
	np1 := person1{
		"kaka",
		30,
	}
	fmt.Printf("%#v\n", np1)// main.person1{string:"kaka", int:30}
	fmt.Println(np1.string, np1.int)// kaka 30
	// 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个

	// 结构体实例化
	// 只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
	// 结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。

	var p1 person
	p1.name = "Sam"
	p1.age = 20
	p1.city = "china"
	fmt.Printf("p1=%v\n", p1)// p1={Sam china 20}
	fmt.Printf("p1=%#v\n", p1)// p1=main.person{name:"Sam", city:"china", age:20}

	// 匿名结构体
	var user struct{
		name string
		age int
	}
	user.name = "Alita"
	user.age = 18
	fmt.Printf("user=%#v\n", user)// user=struct { name string; age int }{name:"Alita", age:18}



	// 创建指针类型结构体
	// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：
	var p2 = new(person)
	// 在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
	p2.age = 29
	p2.name = "Erickson"
	p2.city = "china beijingshi"

	fmt.Printf("%T{ptr:%p}\n", p2, p2)     // *main.person{ptr:0xc0000541e0}
	fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"Erickson", city:"china beijingshi", age:29}


	// 取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。

	p3 := &person{}
	(*p3).name = "luka"
	(*p3).age = 30
	(*p3).city = "china shanghaishi"
	fmt.Printf("%T{ptr:%p}\n", p3, p3)     // *main.person{ptr:0xc000054240}
	fmt.Printf("p3=%#v\n", p3) // p3=&main.person{name:"luka", city:"china shanghaishi", age:30}


	// 结构体初始化
	var p4 person
	fmt.Printf("p4=%#v\n", p4) // p4=main.person{name:"", city:"", age:0}


	// 使用键值对初始化
	// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

	p5 := person{
	   city: "guangzhoushi",
	   name: "modric",
	   age:  31,
	}
	fmt.Printf("p5=%#v\n", p5) // p5=main.person{name:"modric", city:"guangzhoushi", age:31}


	// 也可以对结构体指针进行键值对初始化，例如：
	p6 := &person{
	   name: "kaka",
	   city: "chongqingshi",
	   age:  35,
	}
	fmt.Printf("p6=%#v\n", p6) // p6=&main.person{name:"kaka", city:"chongqingshi", age:35}


	// 当某些字段没有初始值的时候，该字段可以不写。
	// 此时，没有指定初始值的字段的值就是该字段类型的零值。

	p7 := &person{
	   city: "changshashi",
	}
	fmt.Printf("p7=%#v\n", p7) // p7=&main.person{name:"", city:"changshashi", age:0}



	// 使用值的列表初始化
	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：

	p8 := &person{
	   "Angel di Maria",
	   "hangzhoushi",
	   29,
	}
	fmt.Printf("p8=%#v\n", p8) // p8=&main.person{name:"Angel di Maria", city:"hangzhoushi", age:29}

	// 使用这种格式初始化时，需要注意：
	// 1.必须初始化结构体的所有字段。
	// 2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。


	// 面试题
	m := make(map[string]*student)
	stus := []student{
		{name: "kaka", age: 30},
		{name: "modric", age: 32},
		{name: "luka", age: 32},
	}

	for _, stu := range stus {
		m[stu.name] = &student{
			name: stu.name,
			age: stu.age,
		}
	}

	fmt.Println(m, stus)

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}


}

// 构造函数
// Go语言的结构体没有构造函数，我们可以自己实现。
// 例如，下方的代码就实现了一个person的构造函数。
// 因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}


func (p *person) SetAge(newAge int8) {
	p.age = newAge
}


// 方法和接收者
// Go语言中的方法（Method）是一种作用于特定类型变量的函数。
// 这种特定类型变量叫做接收者（Receiver）。
// 接收者的概念就类似于其他语言中的this或者 self。

// 1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。
// 2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
// 3.方法名、参数列表、返回参数：具体格式与函数定义相同。



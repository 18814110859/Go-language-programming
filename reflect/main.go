package main

/**
反射（reflection）是在Java语言出现后迅速流行起来的一种概念。通过反射，你可以获取对
象类型的详细信息，并可动态操作对象。反射是把双刃剑，功能强大但代码可读性并不理想。若
非必要，我们并不推荐使用反射。
Go语言的反射实现了反射的大部分功能，但没有像Java语言那样内置类型工厂，故而无法做
到像Java那样通过类型字符串创建对象实例。在Java中，你可以读取配置并根据类型名称创建对
应的类型，这是一种常见的编程手法，但在Go语言中这并不被推荐。
反射最常见的使用场景是做对象的序列化（serialization，有时候也叫Marshal & Unmarshal）。
例如，Go语言标准库的encoding/json、encoding/xml、encoding/gob、encoding/binary等包就大量
依赖于反射功能来实现。
*/

import (
	"fmt"
	"reflect"
)

/**
一个 Bird 类型 struct
*/
type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying")
}

func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		//输出结果
		fmt.Printf("%d: %s %s = %v/n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

// 输出结果
// 0: Name string = Sparrow
// 1: LifeExpectance int = 3

package main

import "fmt"

/**
 *
 */

type Container []interface{}

// 将元素加到容器中
func (c *Container) Put (elem interface{}) {
	*c = append(*c, elem)
}

// 从 Container 获取一个元素
func (c *Container) Get() interface{} {
	elem := (*c)[0]
	*c = (*c)[1:]
	return elem
}

func main()  {
	// 有一个通用类型的容器，可以进行 Put(val)和 Get()，注意，这里使用了 interface{}做泛型。
	intContainer := &Container{}
	intContainer.Put(1)
	intContainer.Put(2)
	intContainer.Put(3)
	intContainer.Put(4)
	intContainer.Put(5)
	intContainer.Put(6)

	fmt.Println(intContainer)

	// 对某个变量进行 .(type)的转型操作，它会返回两个值，分别是 variable 和 error。 variable 是被转换好的类型，error 表示如果不能转换类型，则会报错。
	// 但是，在把数据取出来时，因为类型是 interface{} ，所以，你还要做一个转型，只有转型成功，才能进行后续操作
	// 断言一个int类型
	elem, ok := intContainer.Get().(int)
	if !ok {
		fmt.Println("unable assert")
	}

	fmt.Printf("%d(%T)\n", elem, elem)
}
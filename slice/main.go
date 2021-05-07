package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/**
 * 需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。
 * 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
 * 2. 切片的长度可以改变，因此，切片是一个可变的数组。
 * 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
 * 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
 * 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
 * 6. 如果 slice == nil，那么 len、cap 结果都等于 0。
 */

func main() {
	arr1 := [...]int{1, 2, 3, 4, 5}
	slice2 := arr1[2:]
	fmt.Printf("arr1: %v length: %d cap: %d points: %p\n", arr1, len(arr1), cap(arr1), &arr1)
	fmt.Printf("slice2: %v length: %d cap: %d points: %p{%p}\n", slice2, len(slice2), cap(slice2), slice2, &slice2)

	// make Slice
	var s1 []int
	var s2 []int
	var s4 = make([]int, 0, 0)
	s3 := []int{1}
	s5 := make([]int, 0, 0)
	fmt.Println(s1, s2, s3, s4, s5)

	// new Array
	// var arr [3]int
	// arr[0] = 1
	arr := [10]int{1, 2, 3, 4, 5}
	s1 = arr[1:7] // [low, high)
	s2 = arr[1:5] // [low, high)
	s2 = append(s2, 1)
	s1[5] = 100
	fmt.Printf("arr: %v{%p}\n", arr, &arr)
	fmt.Printf("s1: %v{%p}{%p}\n", s1, s1, &s1)
	fmt.Printf("s2: %v{%p}{%p}\n", s2, s2, &s1)
	// arr: [1 2 3 4 5 1 100 0 0 0]{0xc000148000}
	// s1: [2 3 4 5 1 100]{0xc000148008}{0xc000128080}
	// s2: [2 3 4 5 1]{0xc000148008}{0xc000128080}

	var s6 = make([]int, 10)
	fmt.Println(*&s6)

	s10 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s10, len(s10), cap(s10))
	//s10[2] += 100
	p := &s10[2]
	*p += 100
	fmt.Println(s10, len(s10), cap(s10))

	// 二维 slice
	data := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}
	fmt.Println(data)

	// 创建一个 struct 的 array
	d := [5]struct {
		x int
	}{}

	sli := d[:]
	d[1].x = 100
	sli[2].x = 200

	fmt.Println(d)
	fmt.Printf("%p, %p, %p, %p, %p, %p\n", &d, &d[0], &d[1], &d[2], &d[3], &d[4])

	s11 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s11)
	s12 := append(s11, 1)
	fmt.Printf("%p\n", &s12)

	// 从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。
	// 如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
	// 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。
	// 或初始化足够长的 len 属性，改用索引号进行操作。
	// 及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。

	// slice中cap重新分配规律：
	s13 := make([]int, 0, 1)
	c := cap(s13)

	for i := 0; i < 50; i++ {
		s13 = append(s13, i)
		n := cap(s13)
		if n > c {
			fmt.Printf("%d->%d\n", c, n)
			c = n
		}
	}

	// slice 拷贝
	s21 := []int{1, 2, 3, 4, 5}
	s22 := make([]int, 10)
	fmt.Println(s21, s22)
	copy(s22, s21)
	fmt.Println(s21, s22)

	data1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s31 := data1[5:10]
	s32 := data1[0:5:6]
	fmt.Printf("slice: %v, len : %d, cap : %d\n", data1, len(data1), cap(data1))
	fmt.Printf("slice: %v, len : %d, cap : %d\n", s31, len(s31), cap(s31))
	fmt.Printf("slice: %v, len : %d, cap : %d\n", s32, len(s32), cap(s32))
	fmt.Println(data1, s31, s32)
	copy(s32, s31)
	fmt.Println(data1, s31, s32)

	// 字符串和切片（string and slice）
	// string底层就是一个byte的数组，因此，也可以进行切片操作。
	str := "Hello world"
	strSlice := str[:]
	fmt.Println(strSlice)

	// 改变string中的字符
	b := []byte(strSlice)
	b[6] = 'G'
	b = b[:8]
	fmt.Println(string(b))

	// 另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x

	// 数组or切片转字符串：
	// strings.Replace(strings.Trim(fmt.Sprint(data1), "[]"), " ", "-", -1)

	// 切片是 Go 中的一种基本的数据结构，使用这种结构可以用来管理数据集合。
	// 切片的设计想法是由动态数组概念而来，为了开发者可以更加方便的使一个数据结构可以自动增加和减少。
	// 但是切片本身并不是动态数据或者数组指针。
	// 切片常见的操作有 re slice、append、copy。
	// 与此同时，切片还具有可索引，可迭代的优秀特性。
	// 在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据。
	// 切片的数据结构
	// 切片本身并不是动态数组或者数组指针。它内部实现的数据结构通过指针引用底层数组，设定相关属性将数据读写操作限定在指定的区域内。
	// 切片本身是一个只读对象，其工作机制类似数组指针的一种封装。
	// 切片（slice）是对数组一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。
	// 这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。
	// 需要注意的是，终止索引标识的项不包括在切片内。
	// 切片提供了一个与指向数组的动态窗口。
	// 给定项的切片索引可能比相关数组的相同元素的索引小。
	// 和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个长度可变的数组。
	// Slice 的数据结构定义如下:
	// type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	// }
	// 切片的结构体由3部分构成，Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。
	// cap 总是大于等于 len 的。

	// 如果想从 slice 中得到一块内存地址，可以这样做：
	s100 := make([]byte, 20)
	ptr100 := unsafe.Pointer(&s100[0])
	fmt.Printf("slice pointer: %p, %p\n", &s100[0], ptr100)

	// 从 Go 的内存地址中构造一个 slice。
	var ptr99 unsafe.Pointer
	var s101 = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr99, 10, 100}

	s102 := *(*[]byte)(unsafe.Pointer(&s101))
	fmt.Printf("slice: %v, %v, %p, %p\n", s101, s102, &s101, &s102)

	// 在 Go 的反射中就存在一个与之对应的数据结构 SliceHeader，我们可以用它来构造一个 slice
	var o []byte
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
	sliceHeader.Cap = 10
	sliceHeader.Len = 10
	sliceHeader.Data = uintptr(ptr99)
	fmt.Printf("slice: %v, %v, %p, %p\n", sliceHeader, o, &sliceHeader, &o)

	// nil 切片被用在很多标准库和内置函数中，描述一个不存在的切片的时候，就需要用到 nil 切片。
	// 比如函数在发生异常的时候，返回的切片就是 nil 切片。nil 切片的指针指向 nil。
	// 空切片一般会用来表示一个空的集合。比如数据库查询，一条结果也没有查到，那么就可以返回一个空切片。

	// 空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。
	// 不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的。

	// 切片扩容
	// 当一个切片的容量满了，就需要扩容了。怎么扩，策略是什么？
	// 主要需要关注的有两点，一个是扩容时候的策略，还有一个就是扩容是生成全新的内存地址还是在原来的地址后追加。

	// 扩容策略
	//slice := []int{10, 20, 30, 40}
	//newSlice := append(slice, 50)
	//fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	//fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	//newSlice[1] += 10
	//fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	//fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

	//array := [4]int{10, 20, 30, 40}
	//slice := array[0:2]
	//newSlice := append(slice, 50)
	//fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	//fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	//newSlice[1] += 10
	//fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	//fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	//fmt.Printf("After array = %v\n", array)

	// Go 中切片扩容的策略：
	// 如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
	// 一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。
	// 注意：扩容扩大的容量都是针对原来的容量而言的，而不是针对原来数组的长度而言的。

	// 切片拷贝
	//array := []int{10, 20, 30, 40}
	//slice := make([]int, 6)
	//n := copy(slice, array)// slicecopy(to, fm slice) int
	//fmt.Println(n,slice, array)

	//slice := make([]byte, 6)
	//n := copy(slice, "abcdef") // slicestringcopy(to []byte, fm string) int
	//fmt.Println(n, slice)


	// 如果用 range 的方式去遍历一个切片，拿到的 Value 其实是切片里面的值拷贝。所以每次打印 Value 的地址都不变。
	// 由于 Value 是值拷贝的，并非引用传递，所以直接改 Value 是达不到更改原切片值的目的的，需要通过 &slice[index] 获取真实的地址
	slice := []int{10, 20, 30, 40}
	fmt.Printf("slice = %v, slice-ptr = %p\n", slice, slice)
	for index, value := range slice {
		fmt.Printf("value = %d , value-addr = %p , slice-addr = %p\n", value, &value, &slice[index])
		*&slice[index] += 10
	}

	fmt.Printf("slice = %v, slice-ptr = %p\n", slice, slice)
}

// 创建切片
// make 函数允许在运行期动态指定数组长度，绕开了数组类型必须使用编译期常量的限制。
// 创建切片有两种形式，make 创建切片，空切片。
// make 和切片字面量

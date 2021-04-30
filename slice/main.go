package main

import "fmt"

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

	//var s6 = make([]int, 100)
	//fmt.Println(*&s6)
	
	s10 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s10, len(s10), cap(s10))
}

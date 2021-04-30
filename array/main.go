package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Array 初始化的几种方式
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]int{1, 2, 4}
var arr4 = [5]int{0: 1, 4: 5}
var arr9 = [2]struct {
	name string
	age  uint
}{
	{"name1", 18},
	{"name2", 28},
}

// 多唯数组 初始化的几种方式
var arr11 [5][3]int
var arr12 = [3][3]int{{1, 2, 3}, {1, 2, 3}}
var arr13 = [...][3]int{{1, 2, 3}, {1, 2, 3}}

func main() {

	// Array 初始化的几种方式
	arr5 := [5]int{1, 2}
	arr6 := [...]int{1, 2}
	arr7 := [5]int{0: 10, 4: 11}
	arr8 := [2]struct {
		name string
		age  uint
	}{
		{"name1", 18},
		{"name2", 28},
	}

	// 多唯数组 初始化的几种方式
	arr14 := [3][3]int{{1, 2, 3}, {1, 2, 3}}
	arr15 := [...][4]int{{1, 2, 3}, {1, 2, 3}}

	// 多维数组的遍历
	for k1, v1 := range arr14 {
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d):%d\n", k1, k2, v2)
		}
	}

	fmt.Println()
	fmt.Println(arr2, arr3, arr4, arr5, arr6, arr7, arr8, arr9)
	fmt.Println(arr11, arr12, arr13, arr14, arr15)

	// 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

	printArr(&[5]int{1, 2, 3, 4, 5})

	rand.Seed(time.Now().Unix())
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}

	fmt.Println(b)
	fmt.Println(sumArr(b))

	findArr([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, 8)
}

func sumArr(s [10]int) int {
	var sum int
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d{%p}\n", s[i], &s[i])
		sum += s[i]
	}

	for _, v := range s {
		fmt.Printf("%d{%p}\n", v, &v)
	}

	return sum
}

func findArr(arr [10]int, target int) {
	for i := 0; i < len(arr); i++ {
		other := target - arr[i]
		for j := i + 1; j < len(arr); j++ {
			if other == arr[j] {
				fmt.Printf("{%d:%d, %d:%d}", i, arr[i], j, arr[j])
			}
		}
	}
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

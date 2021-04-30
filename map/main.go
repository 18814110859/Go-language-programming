package main

import (
	"fmt"
	"strings"
)

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray  []string
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}


func MapStrToInt(arr []string, fn func(s string) int) []int {
	var newArray []int

	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}

	return newArray
}


func Filter(arr []int, fn func(n int) bool) []int {
	var newArray []int
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}


func Reduce(arr []string, fn func(s string) int) int {
	sum := 0
	for _, it := range arr {
		sum += fn(it)
	}
	return sum
}





func main() {
	var list = []string{"Hao", "Chen", "MegaEase"}

	x := MapStrToStr(list, func(s string) string {
		return strings.ToUpper(s)
	})
	// 输出：[HAO CHEN MEGAEASE]
	fmt.Printf("%v\n", x)

	y := MapStrToInt(list, func(s string) int {
		return len(s)
	})
	// 输出：[3 4 8]
	fmt.Printf("%v\n", y)

	z := Reduce(list, func(s string) int {
		return len(s)
	})
	// 输出：15
	fmt.Printf("%v\n", z)

	var intSet = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// intSet = append(intSet, 11)
	input := Filter(intSet, func(n int) bool {
		return n%2 == 1
	})
	// 输出：[1 3 5 7 9]
	fmt.Printf("%v\n", input)

	input1 := Filter(intSet, func(n int) bool {
		return n > 5
	})
	// 输出：[6 7 8 9 10]
	fmt.Printf("%v\n", input1)
}





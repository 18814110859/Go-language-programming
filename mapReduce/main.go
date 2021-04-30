package main

import (
	"fmt"
	"reflect"
	"strings"
)

/**
 * 不做任何类型检查的简单版的泛形
 */
func Map(data interface{}, fn interface{}) []interface{} {
	// reflect是 Go 的反射机制包，作用是在运行时检查类型。
	vFn := reflect.ValueOf(fn)
	vData := reflect.ValueOf(data)
	result := make([]interface{}, vData.Len())

	for i := 0; i < vData.Len(); i++ {
		result[i] = vFn.Call([]reflect.Value{vData.Index(i)})[0].Interface()
	}

	return result
}

func main() {
	square := func(x int) int {
		return x * x
	}

	nums := []int{1, 2, 3, 4}

	squaredArr := Map(nums, square)
	fmt.Println(squaredArr)

	upCase := func(s string) string {
		return strings.ToUpper(s)
	}

	str := []string{"Hap", "Chen", "MegaEase"}
	upStr := Map(str, upCase)
	fmt.Println(upStr)

	t := Map([]int{5}, 5)
	fmt.Println(t)

}


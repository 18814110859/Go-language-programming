package main

import (
	"fmt"
	"reflect"
)

/**
 * 使用反射来做这些东西会有一个问题，那就是代码的性能会很差。所以，上面的代码不能用在需要高性能的地方。
 */

// 员工信息
type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

func main() {
	intList := []int{1, 2, 3, 4, 5, 6}
	list := []string{"1", "2", "3", "4", "5", "6"}

	fmt.Println(Transform(intList, func(x int) int {
		return x * x
	}))

	fmt.Println(Transform(list, func(x string) string {
		return x + x + x
	}))

	var employeeList = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 8000},
		{"Alice", 34, 10, 5000},
		{"Jack", 34, 0, 6000},
		{"Marry", 34, 5, 4000},
		{"Mike", 34, 10, 8000},
	}

	fmt.Println(TransformInPlace(employeeList, func(e Employee) Employee {
		e.Salary += 1000
		e.Age += 1
		return e
	}))

	fmt.Println(Filter(employeeList, func(e Employee) bool {
		return e.Salary > 6000
	}))

}


/**
 * 健壮版的 Generic Reduce
 */
func Reduce(slice, pairFunc, zero interface{}) interface{} {
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("reduce: wrong type, not slice")
	}

	len := sliceInType.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sliceInType.Index(0)
	}

	fn := reflect.ValueOf(pairFunc)
	elemType := sliceInType.Type().Elem()

	if !verifyFuncSignature(fn, elemType, elemType, elemType) {
		t := elemType.String()
		panic("reduce: function must be of type func(" + t + ", " + t + ") " + t)
	}

	var ins [2]reflect.Value
	ins[0] = sliceInType.Index(0)
	ins[1] = sliceInType.Index(1)

	out := fn.Call(ins[:])[0]

	for i := 2; i < len; i++ {
		ins[0] = out
		ins[1] = sliceInType.Index(i)
		out = fn.Call(ins[:])[0]
	}

	return out.Interface()
}


/**
 * 健壮版的 Generic Filter
 */
func Filter(slice, function interface{}) interface{} {
	result, _ := filter(slice, function, false)
	return result
}

func FilterInPlace(slicePtr, function interface{}) {
	in := reflect.ValueOf(slicePtr)
	if in.Kind() == reflect.Ptr {
		panic("FilterInPlace: wrong type, not a pointer to slice")
	}
	_, n := filter(in.Elem().Interface(), function, true)
	in.Elem().SetLen(n)
}

var boolType = reflect.ValueOf(true).Type()

/**
 * 健壮版的 Generic Filter
 */
func filter(slice, function interface{}, inPlace bool) (interface{}, int) {
	// 检查 slice 是否是 slice
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("filter: wrong type, not a slice")
	}

	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, boolType) {

	}

	var which []int
	for i := 0; i < sliceInType.Len(); i++ {
		if fn.Call([]reflect.Value{sliceInType.Index(i)})[0].Bool() {
			which = append(which, i)
		}
	}

	out := sliceInType
	if !inPlace {
		out = reflect.MakeSlice(sliceInType.Type(), len(which), len(which))
	}

	for i := range which {
		out.Index(i).Set(sliceInType.Index(which[i]))
	}

	return out.Interface(), len(which)
}




/**
 * 健壮版的 Generic Map
 * 代码中没有使用 Map 函数，因为和数据结构有含义冲突的问题，所以使用Transform，这个来源于 C++ STL 库中的命名。
 * 有两个版本的函数，一个是返回一个全新的数组 Transform()，一个是“就地完成” TransformInPlace()。
 * 在主函数中，用 Kind() 方法检查了数据类型是不是 Slice，函数类型是不是 Func。
 * 检查函数的参数和返回类型是通过 verifyFuncSignature() 来完成的：NumIn()用来检查函数的“入参”；NumOut() ：用来检查函数的“返回值”。
 * 如果需要新生成一个 Slice，会使用 reflect.MakeSlice() 来完成。
 */
func Transform(slice, fn interface{}) interface{} {
	return transform(slice, fn, false)
}

func TransformInPlace(slice, fn interface{}) interface{} {
	return transform(slice, fn, true)
}

func transform(slice, function interface{}, inPlace bool) interface{} {
	// 检查切片类型是否为切片
	sliceInType := reflect.ValueOf(slice)
	if sliceInType.Kind() != reflect.Slice {
		panic("transform: not slice")
	}

	// 检查func 是否为func
	fn := reflect.ValueOf(function)
	elemType := sliceInType.Type().Elem()
	if !verifyFuncSignature(fn, elemType, nil) {
		panic("trasform: function must be of type func(" + sliceInType.Type().Elem().String() + ") outputElemType")
	}

	sliceOutType := sliceInType
	if !inPlace {
		sliceOutType = reflect.MakeSlice(reflect.SliceOf(fn.Type().Out(0)), sliceInType.Len(), sliceInType.Len())
	}
	for i := 0; i < sliceInType.Len(); i++ {
		sliceOutType.Index(i).Set(fn.Call([]reflect.Value{sliceInType.Index(i)})[0])
	}

	return sliceOutType.Interface()
}

func verifyFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}

	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}

	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

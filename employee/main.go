package main

import (
	"fmt"
)

// 员工信息
type Employee struct {
	Name		string
	Age			int
	Vacation	int
	Salary		int
}

/**
 * 员工统计func
 */
func EmployeeCountIf(list []Employee, fn func(e *Employee) bool) int {
	count := 0
	for i, _ := range list {
		if fn(&list[i]) {
			count += 1
		}
	}
	return count
}

/**
 * 员工过滤func
 */
func EmployeeFilterIn(list []Employee, fn func(e *Employee) bool) []Employee {
	var newEmployee []Employee
	for i, _ := range list {
		if fn(&list[i]) {
			newEmployee = append(newEmployee, list[i])
		}
	}

	return newEmployee
}

/**
 * sum 员工数据
 */
func EmployeeSumIf(list []Employee, fn func(e *Employee) int) int {
	var sum = 0
	for i, _ := range list {
		sum += fn(&list[i])
	}

	return sum
}




func main() {
	var employeeList = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 8000},
		{"Alice", 34, 10, 5000},
		{"Jack", 34, 0, 6000},
		{"Marry", 34, 5, 4000},
		{"Mike", 34, 10, 8000},
	}

	// 统计有多少员工大于40岁
	old := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Age > 40
	})

	fmt.Printf("old people: %d\n", old)

	// 统计员工薪水大于6000的
	highPay := EmployeeCountIf(employeeList, func(e *Employee) bool {
		return e.Salary > 6000
	})

	fmt.Printf("High Salary people: %d\n", highPay)

	// 过滤出没有假期的员工
	notVacation := EmployeeFilterIn(employeeList, func(e *Employee) bool {
		return e.Vacation == 0
	})

	fmt.Printf("not vacction people: %v\n", notVacation)

	// 统计所有员工的薪资总和
	totalPay := EmployeeSumIf(employeeList, func(e *Employee) int {
		return e.Salary
	})

	fmt.Printf("total salary: %d\n", totalPay)

	// 统计40岁以下员工的薪资总和
	youngerPay := EmployeeSumIf(employeeList, func(e *Employee) int {
		if e.Age < 40 {
			return e.Salary
		}
		return 0
	})

	fmt.Printf("total younger salary: %d\n", youngerPay)
}

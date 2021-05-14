package main

import "fmt"

type Dog struct {
	name string
}

func (d *Dog) Move() {
	fmt.Println("狗会动")
}

func (d *Dog) Say() {
	fmt.Println("汪汪汪")
}

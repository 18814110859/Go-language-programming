package main

import "fmt"

type Cat struct {
	name string
}

func (c *Cat) Say() {
	fmt.Println("喵喵喵")
}

func (c *Cat) Move() {
	fmt.Println("猫会动")
}

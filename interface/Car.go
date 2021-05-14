package main

import "fmt"

type Car struct {
	brand string
}

func (c *Car) Move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

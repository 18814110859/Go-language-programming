package main

import "fmt"

type Dryer struct {
}

func (d *Dryer) Dry() {
	fmt.Println("甩一甩")
}

func (d *Dryer) Wash() {
	fmt.Println("洗刷刷")
}

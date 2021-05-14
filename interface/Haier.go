package main

import "fmt"

type Haier struct {
	name string
	Dryer
}

func (h *Haier) Wash() {
	fmt.Println(h.name, "洗刷刷")
}

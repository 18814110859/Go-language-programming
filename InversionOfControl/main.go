package main

import "fmt"

type Widget struct {
	X int
	Y int
}

type Label struct {
	Widget
	Text string
}

//label := Label{Widget{10, 10}, "State"}
//label.x = 11
//label.y = 12

type Button struct {
	Label
}

type ListBox struct {
	Widget
	Text []string
	Index int
}


type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

// 对于 Lable 来说，只有 Painter ，没有Clicker；对于 Button 和 ListBox来说，Painter 和Clicker都有
func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//func (label Label) Click() {
//	fmt.Printf("%p:Label.Click(%q)\n", &label, label.Text)
//}

func (button Button) Paint() {
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

func (box ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%s)\n", box.Text)
}

func(box ListBox) Click() {
	fmt.Printf("ListBox.Click(%s)\n", box.Text)
}

func NewButton(x int, y int, text string) Button {
	return Button{Label{Widget{x, y}, text}}
}

func main() {
	label := Label{Widget{10, 10}, "State:"}
	button1 := Button{Label{Widget{10, 10}, "OK"}}
	button2 := NewButton(50, 60, "Cancel")
	listBox := ListBox{Widget{10, 20}, []string{"AB", "CD", "EF", "GH"}, 0}

	for _, painter := range []Painter{label, button1, button2, listBox} {
		painter.Paint()
	}


	for _, widget := range []interface{}{label, button1, button2, listBox} {
		widget.(Painter).Paint()
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println()
	}
}







package main

import "fmt"

type User struct {
	Name  string
	Email string
}

type Master struct {
	User
	Title string
}

type Manager struct {
	*User
	Title string
}

func (user *User) Test() {
	fmt.Printf("%p, %v\n", user, user)
}

func (user *User) ToString() string {
	return fmt.Sprintf("user %p, %v", user, user)
}

func (user *User) Notify() {
	fmt.Printf("%v : %v prt:%p\n", user.Name, user.Email, user)
}

func (manager *Manager) ToString() string {
	return fmt.Sprintf("manager %p, %v", manager, manager)
}

// 普通函数与方法的区别
// 1.对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
// 2.对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以。

func main() {
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()

	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()

	u2.Name = "java"
	u2.Email = "https://docs.oracle.com/en/java"
	u3.Notify()
	u2.Notify()

	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)
	d.ValueTest()
	p.ValueTest()

	d.PointerTest()
	p.PointerTest()

	m := Manager{&User{"Tom", "tom@email.com"}, "Administrator"}
	m1 := &m

	fmt.Printf("Manager: %v, %v, %p\n", m, &m, &m)
	//fmt.Println(m.User.ToString())
	//fmt.Println(m.ToString())
	fmt.Printf("Manager1: %v, %v, %p\n", m1, &m1, &m1)

	m2 := Master{User{"Sum", "Sum@email.com"}, "Master"}
	m3 := &m2
	fmt.Printf("Master: %v, %v, %p\n", m2, &m2, &m2)
	fmt.Printf("Master1: %v, %p, %p\n", m3, &m3, m3)

	// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
	// instance.method(args...) ---> <type>.func(instance, args...)
	fmt.Println("Golang 表达式")
	u4 := User{"Sum", "Sum@email.com"}
	u4.Test()

	mValue := u4.Test
	mValue2 := (*User).Test
	mValue()
	mValue2(&u4)

	mPointer := (*User).Test
	mPointer(&u4)

	mv := User.TestValue
	mv(u4)
	mp := (*User).TestPointer
	mp(&u4)
	// *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2 := (*User).TestValue
	mp2(&u4)

	var p1 *Data = nil
	p1.PointerTest()
	(*Data)(nil).PointerTest()

	// p1.ValueTest() //  invalid memory address or nil pointer dereference
	// (Data)(nil).ValueTest() // cannot convert nil to type Data
	// Data.ValueTest(nil) // cannot use nil as type Data in argument to Data.ValueTest
	// (*Data).ValueTest(nil) // value method main.Data.ValueTest called using nil *Data pointer

}

type Data struct {
	x int
}

func (d Data) ValueTest() { // func ValueTest(d Data);
	fmt.Printf("Value: %p\n", &d)
}

func (d *Data) PointerTest() { // func PointerTest(d *Data);
	fmt.Printf("Value: %v Pointer: %p{%p}\n", d, d, &d)
}

func (user *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", user, user)
}

func (user User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &user, user)
}

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

type Pperson struct {
	N string
	A int
	S string
	H int
	W int
}

/*
	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
		函数体
	}
*/

// goland中结构体实例是独立的，结构体是值类型，在方法调用中，接收者虽然名为指针类型，但仍然是值拷贝，只不过拷贝的是指针。
// 也就是说，接收者内部的修改不会影响到外部。
// 要修改结构体的值，需要使用指针类型作为接收者。
// 指针类型作为接收者可以修改结构体的值。
// 指针类型作为接收者可以避免复制整个结构体的值，从而提高性能。
func (p Pperson) PrintInfo() {
	fmt.Printf("姓名：%s 年龄：%d 性别：%s 身高：%d 体重：%d\n", p.N, p.A, p.S, p.H, p.W)
}

func (p *Pperson) SetInfo(name string, age int) {
	p.N = name
	p.A = age
}

type myInt int

func (m myInt) PrintInfo() {
	fmt.Println("自定义类型里面的自定义方法")
}

func main() {
	p1 := Person{
		Name: "哈哈",
		Age:  18,
		Sex:  "男",
	}

	p2 := p1
	p2.Name = "哈哈哈哈"
	fmt.Printf("p1: %v, p2: %v\n", p1, p2) // 结构体是值类型
	// 结构体是值类型，当结构体作为参数传递时，会复制整个结构体的值，而不是传递结构体的指针。
	// 这意味着在函数内部对结构体的修改不会影响到原始结构体。
	// 如果你需要在函数内部修改结构体的值，可以使用指针类型作为参数。
	// 指针类型作为参数可以避免复制整个结构体的值，从而提高性能。
	// 指针类型作为参数可以修改原始结构体的值。
	// 指针类型作为参数可以避免复制整个结构体的值，从而提高性能。

	p3 := Pperson{
		N: "嘻嘻",
		A: 18,
		S: "男",
		H: 170,
		W: 60,
	}
	p3.SetInfo("哈吉米", 1338)
	p3.PrintInfo()
	p4 := &Pperson{
		N: "哈哈xixi",
		A: 1338,
		S: "男",
		H: 13270,
		W: 6330,
	}
	p4.PrintInfo()

	var m myInt = 10
	m.PrintInfo()
}

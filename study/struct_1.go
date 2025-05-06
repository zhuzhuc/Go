package main

import "fmt"

type (
	myInt int // 自定义类型
	// type myFn func(int, int) int // 自定义函数类型
	myFloat = float64 // 自定义类型别名
)

// 结构体
type person struct {
	name string
	age  int
	sex  string
} // 结构体首字母小写，表示私有只能在当前包中使用，结构体首字母大写，表示公有，可以在其他包中使用

func main() {
	var a myInt = 1
	fmt.Printf("a的类型是%T\n", a) //	var b int = 1
	var b myFloat = 1.1
	fmt.Printf("b的类型是%T\n", b) //	var c float64 = 1.1
	// 结构体
	var p person // 实例化Person结构体
	p.name = "Tom"
	p.age = 18
	p.sex = "男"
	fmt.Printf("值是%v 类型：%T\n", p, p)
	fmt.Printf("值是%#v 类型：%T\n", p, p)

	p2 := new(person) // 实例化Person结构体
	p2.name = "Alice"
	p2.age = 22
	(*p2).sex = "女"
	fmt.Printf("值是%v 类型：%T\n", p2, p2)
	fmt.Printf("值是%#v 类型：%T\n", p2, p2)

	p3 := &person{}
	p3.name = "Bob"
	p3.age = 25
	p3.sex = "男"
	fmt.Printf("值是%v 类型：%T\n", p3, p3)
	fmt.Printf("值是%#v 类型：%T\n", p3, p3)

	p4 := person{
		name: "Tomhfdfgf",
		age:  99,
		sex:  "男",
	}
	fmt.Printf("值是%v 类型：%T\n", p4, p4)
	fmt.Printf("值是%#v 类型：%T\n", p4, p4)

	p5 := &person{
		name: "GOGOGO",
		age:  18,
		sex:  "男",
	}
	fmt.Printf("值是%v 类型：%T\n", p5, p5)
	fmt.Printf("值是%#v 类型：%T\n", p5, p5)

	p6 := person{
		name: "MOMOMO",
	}
	fmt.Printf("值是%v 类型：%T\n", p6, p6)
	fmt.Printf("值是%#v 类型：%T\n", p6, p6)
	p7 := &person{
		"PPOPOPO",
		999,
		"男",
	}
	fmt.Printf("值是%v 类型：%T\n", p7, p7)
	fmt.Printf("值是%#v 类型：%T\n", p7, p7)
}

// 注意在Goland中支持对结构体指针直接使用，来访问结构体的成员，而不需要使用箭头符号。
// 这是因为Goland提供了一种称为“自动解引用”的功能，它可以在访问结构体指针的成员时自动将指针解引用为结构体。
// 这使得代码更加简洁和易读。
// 例如，在上面的代码中，我们可以直接使用p.name来访问结构体p的name成员，而不需要使用p.name来访问结构体p的name成员。
// 这是因为Goland会自动将p解引用为结构体，从而访问结构体的成员。

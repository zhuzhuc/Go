package main

import "fmt"

/*
	1. 接口是一种类型，是一种抽象的类型
	2. 接口是一种规范，是一种约束
	3. 接口是一种抽象的类型，是一种规范，是一种约束
*/
/*
type 接口名 interface {
	方法名1(参数列表1) 返回值列表1
	方法名2(参数列表2) 返回值列表2
	...
}
*/
type Usber interface {
	Start(string, string) string
	Stop()
}

// 如果接口里面有方法的话，必须通过结构体或自定义类型实现接口
type Phone struct {
	Name string
}

// Start 方法，使其参数列表与 Usber 接口定义一致
func (p Phone) Start(arg1, arg2 string) string {
	fmt.Println(p.Name + " start with args: " + arg1 + ", " + arg2)
	return p.Name + " start with args: " + arg1 + ", " + arg2
}

func (p Phone) Stop() {
	fmt.Println(p.Name + " stop")
}

type Camera struct {
	Name string
}

// 修改 Start 方法，使其参数列表与 Usber 接口定义一致
func (c Camera) Start(arg1, arg2 string) string {
	fmt.Println(c.Name + " start with args: " + arg1 + ", " + arg2)
	return c.Name + " start with args: " + arg1 + ", " + arg2
}

func (c Camera) Stop() {
	fmt.Println(c.Name + " stop")
}

func (c Camera) PrintInfo() {
	fmt.Println(c.Name + " print info")
}

func main() {
	p := Phone{Name: "iPhone"}
	c := Camera{Name: "Canon"}
	// 调用 Start 方法时传入参数
	p.Start("arg1", "arg2")
	c.Start("arg1", "arg2")

	var p1 Usber // 接口就是一个数据类型
	p1 = p
	p1.Stop()

	// 直接使用已有的 c 变量
	var c1 Usber = c
	c1.Start("arg1", "arg2")

	c.PrintInfo()
}

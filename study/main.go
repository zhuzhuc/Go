package main

import "fmt"

// 定义变量 var
// a := 1  自动推导类型
// var a int = 1  定义变量并赋值
// var a int  定义变量不赋值
// var a int = 1, b int = 2  定义多个变量并赋值
// var a, b int = 1, 2  定义多个变量并赋值
// var a, b = 1, 2  定义多个变量并赋值
// var a, b = 1, "2"  定义多个变量并赋值

func getUserinfo() (string, int) {
	return "zzc", 20
}

func main() {
	// fmt.Println("Hello, World!")
	fmt.Print("Hello World!")
	// fmt.Printf("Hello, World!")
	fmt.Println("A", "B", "C")
	// print prinln printf区别
	fmt.Println("----------------------------------")
	a := "aaaa" // go语言中变量定义以后必须使用
	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Println("----------------------------------")
	d := 1
	b := 2
	c := 3
	fmt.Println("d=", d, "b=", b, "c=", c)
	fmt.Printf("d=%v b=%v c=%v\n", d, b, c)
	fmt.Printf("d=%v d的类型是%T\n", d, d)
	//print 输出到控制台
	//println 输出到控制台并换行
	//printf 输出到控制台并格式化
	//%s 字符串
	//%d 数字
	//%f 浮点数
	//%t 布尔值
	//%c 字符
	//%v 任意类型
	//%T 任意类型的类型
	//%p 指针
	//%b 二进制
	//%o 八进制
	//%x 十六进制
	//%X 十六进制大写
	//%U 字符的Unicode码点
	//%q 字符的单引号括起来的字符字面值，由Go语法安全地转义
	//%e 科学计数法
	//%E 科学计数法
	//%g 科学计数法或浮点数
	//%G 科学计数法或浮点数
	//%p 指针
	//%v 任意类型
	//%T 任意类型的类型
	//%p 指针
	//%b 二进制
	//%o 八进制
	//%x 十六进制
	//%X 十六进制大写
	//%U 字符的Unicode码点
	//%q 字符的单引号括起来的字符字面值，由Go语法安全地转义
	//%e 科学计数法
	//%E 科学计数法
	//%g 科学计数法或浮点数

	fmt.Println("----------------------------------")
	fmt.Println("Go变量 常量申明 变量命名规则")
	// var 声明变量
	var username string = "zz"
	fmt.Println(username)

	m_ := "zzc"
	fmt.Println(m_)

	var username1 string
	username1 = "112"
	fmt.Println(username1)

	name := "zzc"
	age := 20
	sex := "man"
	fmt.Println(name, age, sex)

	var (
		name1 string
		age1  int
		sex1  string
	)
	name1 = "zzc"
	age1 = 20
	sex1 = "man"
	fmt.Println(name1, age1, sex1)
	name2, age2, sex2 := "zzc", 20, "man"
	fmt.Println(name2, age2, sex2)
	fmt.Println("----------------------------------")
	fmt.Println("Go常量")

	username2, age2 := getUserinfo()
	fmt.Println(username2, age2)

	username3, _ := getUserinfo()
	fmt.Println(username3)
	_, age3 := getUserinfo()
	fmt.Println(age3)
	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
}

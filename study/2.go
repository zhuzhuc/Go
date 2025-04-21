package main

import "fmt"

func main() {
	username := "zzc"
	username = "zzc2"
	username = "zzc3"
	fmt.Println(username)

	// 常量
	// 常量的值不可以改变
	const PI = 3.1415926
	fmt.Println(PI)
	// 声明多个常量
	const (
		username1 = "zzc"
		username2 = "zzc2"
		username3 = "zzc3"
	)
	fmt.Println(username1, username2, username3)

	// 声明多个常量，值相同
	const (
		username4 = "zzc"
		username5
		username6
	)
	fmt.Println(username4, username5, username6)
	const a = iota // a=0,iota是一个计数器，每次调用iota，值加1，初始为0
	// 声明多个常量，值相同，iota每次调用，值加1
	const (
		b = iota // b=0
		c        // c=1
	)
	fmt.Println(a, b, c)
	fmt.Println("---------------------------")
	const (
		n1 = iota // n1=0
		n2        // n2=1
		_         // 跳过
		n3        // n3=3
		n4 = 100  // n4=100
		n5        // n5=100
		n6 = iota // n6=6
	)

	fmt.Println(n1, n2, n3, n4, n5, n6)
	fmt.Println("---------------------------")
	// 多个itoa
	const (
		a1, a2 = iota + 1, iota + 2 // a1=1,a2=2
		b1, b2 = iota + 1, iota + 2 // b1=2,b2=3
		c1, c2 = iota + 1, iota + 2 // c1=3,c2=4
	)
	fmt.Println(a1, a2, b1, b2, c1, c2)
	fmt.Println("---------------------------")
	// 定义变量
	g1, g2 := 20, 30
	fmt.Println(g1, g2)
	fmt.Println("---------------------------")
}

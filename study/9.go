package main

import (
	"fmt"
)

// 流程控制
func main() {
	fmt.Println("-----------------------------------------")
	var s string = ` 
	_____________________________  
	\____    /\____    /\_   ___ \ 
	  /     /   /     / /    \  \/ 
	 /     /_  /     /_ \     \____
	/_______ \/_______ \ \______  /	
	`
	fmt.Println(s)
	a := 10
	if a >= 10 {
		fmt.Println("a > 10")
	}
	flag := false
	if flag {
		fmt.Println("flag = true")
	} else {
		fmt.Println("flag = false")
	}
	age := 18 // 当前区域全局变量
	if age >= 18 {
		fmt.Println("adult", age)
	}

	if aa := 11; aa >= 18 { // 当前区域局部变量
		// 可以在if条件判断中定义变量，但是只能在if条件判断中使用
		fmt.Println("adult", aa)
	} else {
		fmt.Println("not adult", aa)
	}
	var score int
	fmt.Println("Please enter your score:")
	fmt.Scanln(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 75 && score < 90 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	// if else需要注意的细节
	// 1. if后面的条件表达式不能省略
	// 2. if后面的条件表达式的结果必须是bool类型
	// 3. if后面的条件表达式的结果可以是一个变量
	// 4. if后面的条件表达式的结果可以是一个函数
	// 5. if后面的条件表达式的结果可以是一个表达式
	// 6. if后面的条件表达式的结果可以是一个语句
	// 7. if后面的条件表达式的结果可以是一个代码块

	aaa := 10
	bbb := 20
	var max int
	if aaa > bbb {
		max = aaa
	} else {
		max = bbb
	}
	fmt.Println("max = ", max)

	// switch 语句
	// 1. 语法
	// switch 表达式 {
	// 	case 值1:
	// 		代码块1
	// 	case 值2:
	// 		代码块2
	// 	default:
	// 		代码块3
	// }
	// 2. 表达式
	// 表达式的值可以是任意类型
	// 3. 值
	// 值可以是任意类型
	// 4. 代码块
	// 代码块可以是任意类型
	// 5. default
	// default是可选的
	// 6. case
	// case是可选的
	// 7. break
	// break是可选的
	// 8. default
	// default是可选的
	// 9. 执行流程
	// 9.1 表达式的值和case的值进行比较
	// 9.2 如果相等，则执行对应的代码块
	// 9.3 如果不相等，则继续比较下一个case的值
	// 9.4 如果所有的case都不相等，则执行default对应的代码块

	// for
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	/*
		i := 1
		for i <= 10 {
			fmt.Println(i)
			i++
		}
	*/
	// 死循环
	// for {
	// 	fmt.Println("hello")
	// }
	// 循环嵌套
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

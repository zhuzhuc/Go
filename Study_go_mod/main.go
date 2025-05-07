package main

import (
	"fmt"

	"Study_go_mod/calc"
	"Study_go_mod/tools"

	// 匿名导入包
	_ "Study_go_mod/tools"
)

// 执行顺序：
// 1. 先执行main函数
// 2. 再执行init函数
// 3. 再执行main函数中的代码
// 4. 再执行init函数中的代码
func init() {
	fmt.Println("main init")
}

// main包中的init函数优先于mian函数执行
func main() {
	// 调用calc包中的Add函数
	fmt.Println(calc.Add(1, 2))
	// 调用calc包中的Sub函数
	fmt.Println(calc.Sub(1, 2))
	// 调用calc包中的Mul函数
	fmt.Println(calc.Mul(1, 2))
	// 调用calc包中的Div函数
	fmt.Println(calc.Div(1, 2))

	// 调用calc包中的aaa变量
	// fmt.Println(calc.aaa)// 报错，因为aaa是私有变量，只能在calc包中访问
	fmt.Println(calc.Bbb)

	calc.PrintInfo()
	// 调用tools包中的Mul函数
	fmt.Println(tools.Mul(1, 2))
	// 调用tools包中的PrintInfo函数
	tools.PrintInfo()
	// 调用tools包中的SortIntAsc函数
	fmt.Println(tools.SortIntAsc([]int{1, 3, 4, 9, 10, 3, 4, 5}))
	// 调用tools包中的SortIntDesc函数
	fmt.Println(tools.SortIntDesc([]int{1, 2, 3, 4, 5}))
}

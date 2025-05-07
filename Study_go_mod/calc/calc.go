package calc

import "fmt"

var (
	aaa = "私有变量"
	Bbb = "公有变量"
)

// Add 加法
func Add(a, b int) int {
	return a + b
}

// Sub 减法
func Sub(a, b int) int {
	return a - b
}

// Mul 乘法
func Mul(a, b int) int {
	return a * b
}

// Div 除法
func Div(a, b int) int {
	return a / b
}

func PrintInfo() {
	fmt.Println(aaa)
}

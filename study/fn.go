package main

import "fmt"

func sunFn(a int, b int) int {
	return a + b
}

func subFn(a int, b int) int {
	return a - b
}

func sunFn2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFN3(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func calc2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func fn1(n int) {
	if n > 0 {
		fmt.Printf("n = %d\n", n)
		n--
		fn1(n)
	}
}

func fn2(n int) int {
	// fn1()
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}
}

func main() {
	var a, b int
	fmt.Println("请输入两个数字：")
	fmt.Scanln(&a, &b)
	fmt.Println("a + b =", a+b)
	sum1 := sunFn(a, b)
	fmt.Println("sum1 =", sum1)
	sub1 := subFn(a, b)
	fmt.Println("sub1 =", sub1)
	sum2 := sunFn2(1, 2, 3, 4, 5)
	fmt.Println("sum2 =", sum2)
	sum3 := sumFN3(1, 2, 3, 4, 5)
	fmt.Println("sum3 =", sum3)
	sum4, sub4 := calc(a, b)
	fmt.Println("sum4 =", sum4, "sub4 =", sub4)
	sum5, sub5 := calc2(a, b)
	fmt.Println("sum5 =", sum5, "sub5 =", sub5)

	_, sub6 := calc2(a, b)
	fmt.Println("sub6 =", sub6)
	sub7, _ := calc2(a, b)
	fmt.Println("sub7 =", sub7)

	fn1(10)
	// 函数递归调用
	fmt.Println(fn2(100))
}

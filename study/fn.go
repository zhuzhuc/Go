package main

import "fmt"

// sumFn 计算两个整数的和
func sumFn(a int, b int) int {
	return a + b
}

// subFn 计算两个整数的差
func subFn(a int, b int) int {
	return a - b
}

// sumFn2 计算可变参数整数的和
func sumFn2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// sumFN3 计算第一个整数与可变参数整数的和
func sumFN3(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// calc 同时计算两个整数的和与差
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// calc2 同时计算两个整数的和与差，使用命名返回值
func calc2(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// fn1 递归打印从 n 到 1 的整数
func fn1(n int) {
	if n > 0 {
		fmt.Printf("n = %d\n", n)
		n--
		fn1(n)
	}
}

// fn2 递归计算 1 到 n 的整数和
func fn2(n int) int {
	// fn1()
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}
}

// fn3 递归计算 n 的阶乘
func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}

// 闭包
// 闭包是指一个函数可以访问其外部作用域中的变量。
// 闭包的实现方式是将函数作为返回值返回，同时函数中使用了外部作用域中的变量。
// 闭包的优点是可以访问外部作用域中的变量，并且可以保持变量的值。
// 闭包的缺点是会占用内存，因为闭包中会保留外部作用域中的变量。
func adder() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}
} // 函数里面嵌套一个函数，这个函数可以访问外部函数的变量
func adder2() func(y int) int {
	i := 10
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	var num1, num2 int
	fmt.Println("请输入两个数字：")
	fmt.Scanln(&num1, &num2)
	fmt.Println("num1 + num2 =", num1+num2)
	sum1 := sumFn(num1, num2)
	fmt.Println("sum1 =", sum1)
	sub1 := subFn(num1, num2)
	fmt.Println("sub1 =", sub1)
	sum2 := sumFn2(1, 2, 3, 4, 5)
	fmt.Println("sum2 =", sum2)
	sum3 := sumFN3(1, 2, 3, 4, 5)
	fmt.Println("sum3 =", sum3)
	sum4, sub4 := calc(num1, num2)
	fmt.Println("sum4 =", sum4, "sub4 =", sub4)
	sum5, sub5 := calc2(num1, num2)
	fmt.Println("sum5 =", sum5, "sub5 =", sub5)

	_, sub6 := calc2(num1, num2)
	fmt.Println("sub6 =", sub6)
	sub7, _ := calc2(num1, num2)
	fmt.Println("sub7 =", sub7)

	fn1(10)
	// 函数递归调用
	fmt.Println(fn2(100))
	fmt.Println(fn3(5))

	adderFunc := adder()

	fmt.Println(adderFunc(1))
	fmt.Println(adderFunc(2))
	fmt.Println(adderFunc(3))

	adder2 := adder2()
	fmt.Println(adder2(10))
	fmt.Println(adder2(10))
	fmt.Println(adder2(10))
}

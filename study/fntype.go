// type calculation func(int, int) int
package main

import "fmt"

type calc func(int, int) int // 定义一个calc类型

type myInt int // 定义一个myInt类型
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test")
}

func addd(x, y int) int {
	return x + y
}

func subb(x, y int) int {
	return x - y
}

// 自定义一个方法类型
type calc1 func(int, int) int

func call(x, y int, cb calc1) int {
	return cb(x, y)
}

type caaa func(int, int) int

func do(op string) caaa {
	switch op {
	case "+":
		return addd
	case "-":
		return subb
	case "*":
		return func(x, y int) int {
			return x * y
		}
	case "/":
		return func(x, y int) int {
			return x / y
		}
	default:
		return nil
	}
}

// 匿名函数
func main1() {
	// 匿名函数
	// 匿名函数是指没有函数名的函数
	// 匿名函数可以作为参数传递
	// 匿名函数可以作为返回值
	// 匿名函数可以赋值给变量
	// 匿名函数可以作为闭包
	// 匿名函数可以作为协程
	// 匿名函数可以作为定时器
	// 匿名函数可以作为事件处理器
	// 匿名函数可以作为回调函数
	// 匿名函数可以作为函数类型的变量
	// 匿名函数可以作为函数类型的参数
	// 匿名函数可以作为函数类型的返回值
	// 匿名函数可以作为函数类型的字段
	// 匿名函数可以作为函数类型的方法
	// 匿名函数可以作为函数类型的属性
	// 匿名函数可以作为函数类型的常量
	// 匿名函数可以作为函数类型的枚举
	// 匿名函数可以作为函数类型的结构体
	// 匿名函数可以作为函数类型的类
	// 匿名函数可以作为函数类型的接口
}

func main() {
	var a calc = add
	fmt.Println(a(1, 2))
	fmt.Printf("a的类型是%T\n", a)
	var b calc = sub
	fmt.Println(b(1, 2))
	test()
	f := sub
	fmt.Printf("f的类型是%T\n", f)

	var c myInt = 1
	fmt.Printf("c的类型是%T\n", c)
	var d int = 1
	fmt.Printf("d的类型是%T\n", d)

	fmt.Println(int(c) + d)

	// 修正函数调用
	sum := call(10, 5, addd)
	subResult := call(10, 5, subb)
	fmt.Println(sum)
	fmt.Println(subResult)

	j := call(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Println(j)

	aa := do("+")
	bb := do("-")
	cc := do("*")
	dd := do("/")
	fmt.Println(aa(1, 2))
	fmt.Println(bb(1, 2))
	fmt.Println(cc(1, 2))
	fmt.Println(dd(1, 2))

	func() {
		fmt.Println("匿名函数")
	}()

	fn := func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 3))
}

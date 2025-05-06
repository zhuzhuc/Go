package main

import (
	"errors"
	"fmt"
)

func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

func f3() (r int) {
	defer func() {
		r++
	}()
	return r
}

func f4() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}

func f5() (x int) {
	defer func() {
		x++
	}()
	return 5 // 6
}

func f6() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 5
}

func f7() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 5
}

// defer注册要延迟执行的函数时改函数所有的参数都需要确定其值
func cala(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main1() {
	fmt.Println("main1")
}

func main2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	panic("抛出一个异常值")
}

func x(a int, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()
	return a / b
}

// 异常处理
// 异常处理是指在程序运行过程中出现错误或异常情况时，程序能够捕获并处理这些错误或异常情况，以保证程序的正常运行。
// 异常处理通常包括以下几个方面：
// 1. 捕获异常：程序能够捕获到异常情况，并记录异常信息。
// 2. 处理异常：程序能够根据异常情况进行处理，例如输出错误信息、记录日志、进行回滚操作等。
// 3. 恢复程序：程序能够恢复到正常运行状态，继续执行后续代码。
// 4. 异常处理机制：程序能够提供异常处理机制，例如使用 try-catch 语句、使用 defer 语句等。
// 5. 异常处理策略：程序能够提供异常处理策略，例如忽略异常、重试操作、抛出异常等。
// 6. 异常处理原则：程序能够提供异常处理原则，例如避免使用异常处理机制，使用错误码代替异常处理机制等。
// 7. 异常处理工具：程序能够提供异常处理工具，例如使用异常处理库、使用异常处理框架等。
func readFile(filename string) error {
	if filename == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	// defer 延迟调用

	// defer 语句会将函数推迟到外层函数返回之后执行。
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	// 多个 defer 语句按照后进先出的顺序执行。
	// 这意味着最后一个 defer 语句的函数会最先被调用，第一个 defer 语句的函数会最后被调用。
	// defer 语句经常被用于释放资源，例如关闭文件、解锁互斥锁等。
	// defer 语句也可以用于记录函数的执行时间，例如在函数开始时记录开始时间，在函数结束时记录结束时间。
	// fmt.Println("start")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() {
	// 	fmt.Println("zzc")
	// }()
	// fmt.Println("end")

	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
	fmt.Println(f7())

	// x := 1
	// y := 2
	// defer cala("AA", x, cala("BB", x, y))
	// x = 10
	// defer cala("CC", x, cala("DD", x, y))
	// y = 20
	/*
		BB 1 2 3
		DD 10 2 12
		CC 10 12 22
		AA 1 3 4
	*/
	main1()
	main2()
	fmt.Println(x(10, 0))
	fmt.Println("-----------------------------------------")
	myFn()
	fmt.Println("继续执行...")
}

package main

import "fmt"

func main() {
	var s string = ` 
▗▄▄▄▄▖▗▄▄▄▄▖ ▗▄▄▖
   ▗▞▘   ▗▞▘▐▌   
 ▗▞▘   ▗▞▘  ▐▌   
▐▙▄▄▄▖▐▙▄▄▄▖▝▚▄▄▖
                 
	`
	fmt.Println(s)
	// for range
	// for range 是go语言特有的一种的迭代结构，它可以配合数组、切片、字符串、map以及通道（channel）使用。
	// 格式：
	// for key, value := range collection {
	//     // do something
	// }
	// 1. 数组、切片、字符串
	// 数组、切片、字符串返回索引和值
	// 数组
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}
	str := "zzc"
	for k, v := range str {
		fmt.Println("key=", k, "value=", v)
	}

	arr1 := []string{"java", "python", "c++", "go", "php"}
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
		// if i < len(arr1)-1 {
		// 	fmt.Print(" ")
		// }
	}
	fmt.Println()
	for _, v := range arr1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	// switch case
	// switch case 是go语言特有的一种的条件结构，它可以配合数组、切片、字符串、map以及通道（channel）使用。
	// 格式：
	// switch expression {
	// case value1:
	//     // do something
	// case value2:
	//     // do something
	// default:
	//     // do something
	// }

	var aaa string
	fmt.Println("please enter file name:")
	fmt.Scanln(&aaa)
	switch aaa {
	case ".html":
		fmt.Println("html")
		break
	case ".css":
		fmt.Println("css")
		break
	case ".js":
		fmt.Println("js")
		break
	default:
		fmt.Println("unknown")
	}

	// 穿透
	var bbb int
	fmt.Println("please enter number:")
	fmt.Scanln(&bbb)
	switch bbb {
	case 1:
		fmt.Println("one")
		fallthrough // 穿透，会继续执行下一个case
	case 2:
		fmt.Println("two")
		fallthrough // 穿透，会继续执行下一个case
	case 3:
		fmt.Println("three")
		break // break可写可不写，默认会有break
	default:
		fmt.Println("unknown")
	}

	var n int
	fmt.Println("please enter a number:")
	fmt.Scanln(&n)
	switch {
	case n%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("odd")
	}

	var age int
	fmt.Println("please enter your age:")
	fmt.Scanln(&age)
	switch {
	case age >= 18:
		fmt.Println("adult")
	default:
		fmt.Println("not adult")
	}
	// 类型断言
	var ccc interface{}
	ccc = 10
	value, ok := ccc.(int) // 类型断言
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("type assertion failed")
	}
}

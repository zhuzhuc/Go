package main

import "fmt"
import "unsafe"

func main() {
	// 整型
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Printf("c=%v 类型:%T\n", c, c)
	fmt.Println("----------------------")
	//int8(-128, 127)
	var num int8 = 98;
	fmt.Printf("num=%v 类型:%T\n", num, num)
	//uint(0-255)
	
	var s int8 = 10;
	fmt.Printf("s=%v 类型:%T\n", s, s)
	fmt.Println(unsafe.Sizeof(s))
	//unsafe.Sizeof
	//输出一个字节(1)
	var  p uint8  = 100
	fmt.Printf("p=%v 类型:%T\n", p, p)
	fmt.Println(unsafe.Sizeof(p))
	
	fmt.Println("-----------------------")
	var a1 int32 = 10
	var a2 int64 = 21
	fmt.Println(int64(a1) + a2)
	//高位像低位转换	
	var n1 int16 = 130
	fmt.Println(int8(n1))//需要注意类型的范围
	
	number := 30
	fmt.Printf("num=%v 类型:%T\n", number, number)
	fmt.Println(unsafe.Sizeof(number))
	
}

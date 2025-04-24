package main

import (
	"fmt"
)

func main() {
	// byte 和rune类型
	a := 'a'
	fmt.Printf("a=%v--%T\n", a, a)
	// 原样输出字符a
	fmt.Printf("a=%c\n", a)
	fmt.Println("------------------------------")
	/*
		byte类型
		byte类型是uint8的别名，用于表示ASCII码的一个字符
	*/
	var b byte = 'a'
	fmt.Printf("b=%v--%T\n", b, b)
	fmt.Printf("b=%c\n", b)
	fmt.Println("------------------------------")
	/*
		rune类型
		rune类型是int32的别名，用于表示一个Unicode码点
	*/
	var c rune = '中'
	fmt.Printf("c=%v--%T\n", c, c)
	fmt.Println("------------------------------")
	/*
		字符串的遍历
	*/
	str := "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c)\n", str[i], str[i])
		// fmt.Println(str[i])
	}
	for _, v := range str {
		fmt.Printf("%v(%c)\n", v, v)
	}
	// 一个汉字占3个字节，一个字母占1个字节

	// 字符串修改
	str1 := "hello"
	fmt.Println(str1)
	// str1[0] = 'z'
	// fmt.Println(str1)
	// 字符串修改
	str2 := "hello"
	byteStr := []byte(str2)
	byteStr[0] = 'z'
	// str2 = string(byteStr)
	fmt.Println(string(byteStr))
	str2 = "zzc"
	fmt.Println(str2)

	str5 := "你好 goland"
	runeStr := []rune(str5)
	runeStr[0] = '是'
	fmt.Println(string(runeStr))
}

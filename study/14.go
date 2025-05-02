package main

import "fmt"

func main() {
	// 值类型
	// 1. 基本数据类型 int float string bool
	// 2. 数组和结构体
	// 3. 指针
	// 4. 函数
	// 5. 接口
	// 6. 切片
	// 7. 字典
	// 8. 通道
	// 9. 变量
	// 10. 常量
	// 11. 类型别名
	// 12. 类型定义
	// 13. 类型转换
	// 14. 类型断言
	// 15. 类型推导
	// 16. 类型组合
	// 17. 类型嵌套
	// 18. 类型嵌入
	a := 10
	b := 10
	a = 20
	fmt.Println(a, b)

	a1 := [...]int{1, 2, 3}
	a2 := a1
	fmt.Println(a1)
	fmt.Println(a2)
	a1[0] = 100
	fmt.Println(a1)
	fmt.Println(a2)

	// 引用类型
	// 1. 切片
	// 2. 字典
	// 3. 通道
	// 4. 函数
	// 5. 接口
	// 6. 指针
	// 7. 结构体
	// 8. 数组
	// 9. 变量
	// 10. 常量
	// 切片
	// 1. 切片是数组的引用
	// 2. 切片是一个引用类型
	// 3. 切片是一个可变长度的数组
	/*
		    值类型：改变变量副本的值，不会影响到原始变量的值
			引用类型：改变变量副本的值，会影响到原始变量的值
	*/
	a3 := []int{1, 2, 3, 4, 5}
	a4 := a3
	a3[0] = 100
	a4[0] = 200
	fmt.Println(a3)
	fmt.Println(a4)

	// 多维数组
	// 1. 二维数组
	// 2. 三维数组
	// 3. 四维数组
	// 4. 多维数组
	// 5. 多维数组的遍历
	// 6. 多维数组的切片
	// 7. 多维数组的函数
	// 8. 多维数组的指针
	// 9. 多维数组的引用
	// 10. 多维数组的常量
	var aa [3][4]int
	fmt.Println(aa)
	aa[0][0] = 1
	aa[0][1] = 2
	aa[0][2] = 3
	aa[0][3] = 4
	aa[1][0] = 5
	aa[1][1] = 6
	aa[1][2] = 7
	aa[1][3] = 8
	aa[2][0] = 9
	aa[2][1] = 10
	aa[2][2] = 11
	aa[2][3] = 12

	fmt.Println(aa)

	for i := 0; i < len(aa); i++ {
		for j := 0; j < len(aa[i]); j++ {
			fmt.Printf("%v%v", aa[i][j], " ")
		}
	}
	fmt.Println()

	ass := [3][2]string{
		{"php", "java"},
		{"goland", "rust"},
		{"js", "ts"},
	}

	for i := len(ass) - 1; i >= 0; i-- {
		for j := 0; j < len(ass[i]); j++ {
			fmt.Printf("%v%v", ass[i][j], " ")
		}
	}
	fmt.Println()

	for _, v := range ass {
		for _, vv := range v {
			fmt.Printf("%v%v", vv, " ")
		}
	}
	fmt.Println()
	asss := [...][2]string{
		{"php", "java"},
		{"goland", "rust"},
		{"js", "ts"},
	}
	for i := len(asss) - 1; i >= 0; i-- {
		for j := 0; j < len(asss[i]); j++ {
			fmt.Printf("%v%v", asss[i][j], " ")
		}
	}
	fmt.Println()
}

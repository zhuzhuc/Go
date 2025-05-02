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

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	// var arr2 [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v%v", arr[i], " ")
	}
	fmt.Println()

	// 数组初始化
	// 1. 定义时初始化
	var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	// 2. 定义时初始化，省略长度
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

	var strarr [5]string = [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strarr)
	strarr2 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(strarr2)

	var stRarr [3]string
	stRarr[0] = "php"
	stRarr[1] = "java"
	stRarr[2] = "goland"
	fmt.Println(stRarr)
	fmt.Println("----------------------------------")

	aaa := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(aaa)

	sss := [...]string{"cpp", "goland", "java", "rust"}
	fmt.Println(sss)
	fmt.Println(len(sss))

	as := [...]int{0: 1, 1: 10, 2: 30, 5: 50}
	fmt.Println(len(as))

	ass := [3]int{23, 34, 5}
	for i := 0; i < len(ass); i++ {
		fmt.Println(ass[i])
	}

	abs := [...]string{"php", "java", "nodejs", "goland", "js"}

	for k, v := range abs {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	e := [...]int{12, 23, 45, 67, 2, 5}
	sum := 0
	for i := 0; i < len(e); i++ {
		sum += e[i]
	}
	fmt.Println("数组的和:", sum)
	fmt.Println("数组的平均值:", sum/len(e))

	for v := range e {
		fmt.Printf("%v ", e[v])
	}
	fmt.Println()

	var lenn int
	fmt.Println("请输入数组的长度")
	fmt.Scanln(&lenn)

	// 使用切片替代数组
	arr5 := make([]int, lenn)
	fmt.Println("请输入数组元素，用空格分隔")
	for i := 0; i < lenn; i++ {
		fmt.Scan(&arr5[i])
	}

	// 求出数组的最大值，并得到对应的下标
	maxValue := arr5[0]
	maxIndex := 0
	for i := 1; i < len(arr5); i++ {
		if arr5[i] > maxValue {
			maxValue = arr5[i]
			maxIndex = i
		}
	}

	fmt.Println("输入的数组:", arr5)
	fmt.Println("最大值:", maxValue)
	fmt.Println("最大值的下标:", maxIndex)

	arr6 := [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(arr6); i++ {
		for j := i + 1; j < len(arr6); j++ {
			if arr6[i]+arr6[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
	fmt.Println()
}

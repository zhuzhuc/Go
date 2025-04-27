package main

import "fmt"

func main() {
	// break
	// break语句用于跳出当前循环
	// break语句可以用于for循环、switch语句、select语句等
	// break语句可以用于跳出当前循环，但是不能用于跳出函数

	for i := 1; i <= 10; i++ {
		if i == 9 {
			break // i==9时，跳出循环
		}
		fmt.Printf("%v%v", i, " ")
	}
	fmt.Println()
	// continue
	// continue语句用于跳过当前循环中的剩余语句，然后继续进行下一轮循环
	// continue语句可以用于for循环、switch语句、select语句等
	// continue语句可以用于跳过当前循环中的剩余语句，但是不能用于跳出函数
	for i := 1; i <= 10; i++ {
		if i == 9 {
			continue // i==9时，跳过当前循环中的剩余语句，继续进行下一轮循环
		}
		fmt.Printf("%v%v", i, " ")
	}
	fmt.Println()

label1:
	for i := 1; i <= 10; i++ {
		if i == 9 {
			break label1 // i==9时，跳出label标签
		}
		fmt.Printf("%v%v", i, " ")
	}
	fmt.Println()
	// goto
	// goto语句用于跳转到指定的标签
	// goto语句可以用于跳转到指定的标签，但是不能用于跳出函数
	for i := 1; i <= 10; i++ {
		if i == 9 {
			goto label // i==9时，跳转到label标签
		}
	}
label:
	fmt.Println("goto")

label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2 // j==3时，跳转到label标签
			}
			fmt.Println("i=", i, "j=", j)
		}
	}
}

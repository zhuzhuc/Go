// go中的运算符
package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test")
	return true
}

// /
func main() {
	aa := 6
	bb := 3
	fmt.Println(aa + bb)
	fmt.Println(aa - bb)
	fmt.Println(aa * bb)
	fmt.Println(aa / bb)
	fmt.Println(aa % bb)
	c := aa * bb
	fmt.Println(c)
	ss := 2
	ss++
	fmt.Println(ss)
	fmt.Println("-------------------------------")
	// 比较运算符
	// ==
	// !=
	// >
	// <
	// >=
	// <=
	fmt.Println(aa == bb)
	fmt.Println(aa != bb)
	fmt.Println(aa > bb)
	fmt.Println(aa < bb)
	fmt.Println(aa >= bb)
	fmt.Println(aa <= bb)

	flag := aa > bb
	if flag {
		fmt.Println("aa > bb")
	}
	fmt.Println("-------------------------------")
	// go语言中没有自增自减运算符(包括前置++--)
	// aa++
	// aa--
	// 逻辑运算符
	// && 与
	// || 或
	// ! 非
	// 逻辑运算符的结果为bool类型
	var a bool = true
	var b bool = false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)

	fmt.Println("-------------------------------")
	// 输入aaaab
	fmt.Println("请输入数字")
	var aaaab int
	fmt.Scanln(&aaaab)
	if aaaab > 10 && test() {
		fmt.Println("run")
	}
	// } else {
	// 	fmt.Println("stop")
	// }
	// 赋值运算符
	// =
	// +=
	// -=
	// *=
	// /=
	// %=
	fmt.Println("-------------------------------")
	var aaaaa int = 10
	aaaaa += 10
	fmt.Println(aaaaa)
	aaaaa -= 10
	fmt.Println(aaaaa)
	aaaaa *= 10
	fmt.Println(aaaaa)
	aaaaa /= 10
	fmt.Println(aaaaa)
	aaaaa %= 10
	fmt.Println(aaaaa)

	// 位运算符
	// & 按位与 两位都为1才为1
	// | 按位或  两位都为0才为0，否则为1
	// ^ 按位异或 两位都为1才为0
	// << 左移 高位舍弃，低位补0，乘以2的n次方
	// >> 右移 高位补0，低位舍弃，除以2的n次方
	fmt.Println("-------------------------------")
	var aaaaaa int = 10
	var bbbbbb int = 20
	fmt.Println(aaaaaa & bbbbbb) // 0
	fmt.Println(aaaaaa | bbbbbb) // 30
	fmt.Println(aaaaaa ^ bbbbbb) // 30
	fmt.Println(aaaaaa << 1)     // 20
	fmt.Println(bbbbbb >> 1)     // 10
	fmt.Println(aaaaaa >> 1)     // 5
	// 位运算符的结果为int类型

	// 条件运算符
	// ? :
	fmt.Println("-------------------------------")
	var aaaaaaa int = 10
	var bbbbbbb int = 20
	if aaaaaaa > bbbbbbb {
		fmt.Println("aaaaaaa > bbbbbbb")
	}
	if aaaaaaa < bbbbbbb {
		fmt.Println("aaaaaaa < bbbbbbb")
	}
	if aaaaaaa == bbbbbbb {
		fmt.Println("aaaaaaa == bbbbbbb")
	}
	if aaaaaaa != bbbbbbb {
		fmt.Println("aaaaaaa != bbbbbbb")
	}
	if aaaaaaa >= bbbbbbb {
		fmt.Println("aaaaaaa >= bbbbbbb")
	}
	fmt.Println("-------------------------------")
	as := 10
	bs := 20
	as = as + bs
	bs = as - bs
	as = as - bs
	fmt.Println(as)
	fmt.Println(bs)

	// 假如还有100天放假，问：xx个星期零xx天
	var days int = 100
	var weeks int = days / 7
	var days2 int = days % 7
	fmt.Printf("%d个星期零%d天\n", weeks, days2)

	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100),请求出华氏温度对应的摄氏温度
	var F float32 = 100
	C := (F - 32) / 1.8
	fmt.Printf("%.2f华氏度对应摄氏温度为%.2f\n", F, C)

	fmt.Println("-------------------------------")
}

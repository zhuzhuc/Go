package main

import (
	"fmt"
	"unsafe"

	"github.com/shopspring/decimal"
)

func main() {
	var a float32 = 3.12
	fmt.Printf("a=%v--%f, 类型:%T\n", a, a, a)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println("----------------------")
	var c float64 = 3.1415926535
	fmt.Printf("%v--%f--%.2f\n", c, c, c)
	fmt.Printf("%.4f\n", c)

	f1 := 3.333333333
	fmt.Printf("f1=%v--%T\n", f1, f1)

	fmt.Println("-----------------------------")
	// 科学计数法表述浮点类型
	f2 := 3.12e2
	fmt.Printf("f2=%v--%T\n", f2, f2)

	f3 := 3.12e-2
	fmt.Printf("f3=%v--%T\n", f3, f3)

	fmt.Println("------------------------------")

	var num1 float64 = 1.3
	var num2 float64 = 4.2
	d1 := decimal.NewFromFloat(float64(num1)).Add(decimal.NewFromFloat(num2))
	fmt.Printf("d1=%v--%T\n", d1, d1)
	fmt.Println(d1)

	fmt.Println("------------------------------")
	var num3 float64 = 1.3
	var num4 float64 = 4.2
	d2 := decimal.NewFromFloat(num3).Sub(decimal.NewFromFloat(num4))
	fmt.Printf("d2=%v--%T\n", d2, d2)
	fmt.Println(d2)

	fmt.Println("------------------------------")
	var num5 int = 10
	fmt.Println((float64(num5) * 0.3))

	// float转int
	var num6 float64 = 1.9
	fmt.Println(int(num6))

	// bool类型
	var isOk bool = true
	fmt.Printf("isOk=%v--%T\n", isOk, isOk)

	k := true
	if k {
		fmt.Println("k=true")
	} else {
		fmt.Println("k=false")
	}
}

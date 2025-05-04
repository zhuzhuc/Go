package main

import (
	"fmt"
	"strconv"
)

// 类型转换
// 类型转换的格式为：类型(表达式)
func main() {
	// 数值类型之间的转换
	var a int8 = 20
	var b int16 = 40
	fmt.Println(a + int8(b))

	var ab float32 = 20
	var ba float64 = 40
	fmt.Println(ab + float32(ba))

	var aa float32 = 28.999
	var bb int = 40
	fmt.Println(aa + float32(bb))

	var aaa int8 = 20
	var bbb int16 = 140
	fmt.Println(int16(aaa) + bbb)

	// 其他类型转为string
	// int 为%d float为%f bool为%t byte为%c
	var i int = 20
	var f float64 = 12.456
	var t bool = true
	var bbbb byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("str1= %v  类型：%T\n", str1, str1)
	str2 := fmt.Sprintf("%.2f", f)
	fmt.Printf("str2= %v  类型：%T\n", str2, str2)
	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("str3= %v  类型：%T\n", str3, str3)
	str4 := fmt.Sprintf("%c", bbbb)
	fmt.Printf("str4= %v  类型：%T\n", str4, str4)

	// 通过strconv包进行转换
	var num1 int = 20
	str5 := strconv.FormatInt(int64(num1), 10)
	// fmt.Println(str5)
	fmt.Printf("str5= %v  类型：%T\n", str5, str5)

	var fff float32 = 29.22333
	// 参数1：要转换的数字
	// 参数2：格式
	// 参数3：精度
	// 参数4：返回字符串的长度
	str6 := strconv.FormatFloat(float64(fff), 'f', 2, 64)
	fmt.Printf("str6= %v  类型：%T\n", str6, str6)

	var booool bool = true
	str7 := strconv.FormatBool(booool)
	fmt.Printf("str7= %v  类型：%T\n", str7, str7)

	aaaaa := 'z'
	str8 := strconv.FormatUint(uint64(aaaaa), 10)
	fmt.Printf("str8= %v  类型：%T\n", str8, str8)

	// string转换为其他类型
	// strconv.ParseInt()
	// strconv.ParseFloat()
	// strconv.ParseBool()
	// strconv.ParseUint()
	// 参数1：要转换的字符串
	// 参数2：进制
	// 参数3：返回的位数
	// 参数4：返回的类型
	// 返回值：转换后的数字和错误信息
	var ss string = "12345"
	num2, err := strconv.ParseInt(ss, 10, 64)
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Printf("num2= %v  类型：%T\n", num2, num2)

	var ss1 string = "12345.6789"
	num3, err := strconv.ParseFloat(ss1, 64)
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Printf("num3= %.2f  类型：%T\n", num3, num3)
	var ss2 string = "true"
	num4, err := strconv.ParseBool(ss2)
	if err != nil {
		fmt.Println("转换失败")
	}
	fmt.Printf("num4= %v  类型：%T\n", num4, num4)

	sssss, _ := strconv.ParseBool("zzzzc")
	fmt.Println(sssss)
}

///

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello"
	fmt.Println(str)
	var str1 string = "zzc"
	fmt.Println(str1)
	str2 := "zzc"
	// fmt.Println(str2)
	fmt.Printf("str2=%v--%T\n", str2, str2)

	// 字符串转义符
	fmt.Println("------------------------------")

	a := "this is\n a string"
	fmt.Println(a)

	b := "c:\\zzc\\study\\go"
	fmt.Println(b)
	c := "c:\\zzc\\study\\go\\hello.txt"
	fmt.Println(c)

	fmt.Println("------------------------------")
	/*
		多行字符串
	*/
	d := `
	zzc
	study
	go
	`
	fmt.Println(d)
	fmt.Println("------------------------------")
	f := `
	__   __  _______  _______  ______    _______  _______ 
	|  | |  ||       ||       ||    _ |  |       ||       |
	|  |_|  ||    ___||    ___||   | ||  |   _   ||  _____|
	|       ||   |___ |   |___ |   |_||_ |  | |  || |_____ 
	|       ||    ___||    ___||    __  ||  |_|  ||_____  |
	|   _   ||   |___ |   |___ |   |  | ||       | _____| |
	|__| |__||_______||_______||___|  |_||_______||_______|
	`
	fmt.Println(f)

	/*
		字符串的拼接
	*/
	str3 := "zzc"
	str4 := "study"
	str5 := str3 + str4
	fmt.Println(str5)
	/*
		字符串的长度
	*/
	fmt.Println(len(str5))
	/*
		字符串的遍历
	*/
	for i := 0; i < len(str5); i++ {
		fmt.Printf("%c", str5[i])
	}

	// 字符串常见操作
	fmt.Println("------------------------------")
	str6 := "hello"
	str7 := "zzc"
	fmt.Println(str6 == str7)
	fmt.Println(str6 != str7)
	fmt.Println(str6 > str7)
	fmt.Println(str6 < str7)
	strss := fmt.Sprintf("%s %s", str6, str7)
	fmt.Println(strss)
	fmt.Println("------------------------------")
	/*
		字符串的切片
	*/
	str8 := "zzcstudygo"
	fmt.Println(str8[0:3])
	fmt.Println(str8[3:])
	fmt.Println(str8[:3])
	fmt.Println(str8[:])
	/*
		字符串的分割
	*/
	str9 := "zzc,study,go"
	fmt.Println(str9)

	/*
		字符串的替换
	*/
	str10 := "zzc study go"
	fmt.Println(str10)
	str11 := strings.Replace(str10, "zzc", "zzcstudy", 1)
	fmt.Println(str11)
	/*
		字符串的包含
	*/
	str12 := "zzcstudygo"
	fmt.Println(strings.Contains(str12, "zzc"))
	fmt.Println(strings.Contains(str12, "zzcstudy"))
	fmt.Println(strings.Contains(str12, "zzcstudygo"))
	/*
		字符串的前缀和后缀
	*/
	str13 := "zzcstudygo"
	fmt.Println(strings.HasPrefix(str13, "zzc"))
	fmt.Println(strings.HasSuffix(str13, "go"))
	/*
		字符串的索引
	*/
	str14 := "zzcstudygo"
	fmt.Println(strings.Index(str14, "zzc"))
	fmt.Println(strings.Index(str14, "study"))
	fmt.Println(strings.Index(str14, "go"))
	//查找不到返回-1，否则返回下标
	/*
		字符串的拼接,join把切片中的元素用指定的分隔符拼接成一个字符串
	*/
	str15 := "zzc"
	str16 := "study"
	str17 := "go"
	fmt.Println(strings.Join([]string{str15, str16, str17}, "_-_"))
	/*
		字符串的分割	split把字符串按照指定的分隔符分割成一个切片
	*/
	str18 := "zzc-study-go"
	fmt.Println(strings.Split(str18, "-"))

	sssss := []string{"php", "java", "goland", "rust", "c++"}
	fmt.Println(strings.Join(sssss, "_"))
}

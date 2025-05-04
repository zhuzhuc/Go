package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	userinfo := make(map[string]string)
	userinfo["name"] = "zs"
	userinfo["age"] = "18"
	userinfo["gender"] = "男"
	userinfo["address"] = "北京"
	userinfo["phone"] = "13800138000"
	userinfo["address"] = "北京"
	userinfo["phone"] = "13800138000"
	userinfo["email"] = "EMAIL"
	fmt.Println(userinfo)
	// 遍历map
	for key, value := range userinfo {
		fmt.Println(key, value)
	}

	mminfo := map[string]string{
		"userneme": "zzc",
		"age":      "20",
		"gender":   "男",
	}
	fmt.Println(mminfo)

	// map类型的curd
	// 添加
	mminfo["address"] = "北京"
	fmt.Println(mminfo)
	// 修改
	mminfo["address"] = "上海"
	fmt.Println(mminfo)
	// 删除
	delete(mminfo, "address")
	fmt.Println(mminfo)
	// 查询
	value, ok := mminfo["address"]
	if ok {
		fmt.Println(value)
	}

	usinfo := make([]map[string]string, 3, 3)
	fmt.Println(usinfo[0]) // map 不初始化时，默认值为 nil
	if usinfo[0] == nil {
		// 初始化 usinfo[0] 为 map[string]string 类型
		usinfo[0] = make(map[string]string)
		usinfo[0]["name"] = "zs"
		usinfo[0]["age"] = "18"
		usinfo[0]["gender"] = "男"
	}
	if usinfo[1] == nil {
		// 初始化 usinfo[1] 为 map[string]string 类型
		usinfo[1] = make(map[string]string)
		usinfo[1]["name"] = "ls"
		usinfo[1]["age"] = "20"
		usinfo[1]["gender"] = "女"
	}
	if usinfo[2] == nil {
		// 初始化 usinfo[2] 为 map[string]string 类型
		usinfo[2] = make(map[string]string)
		usinfo[2]["name"] = "ww"
		usinfo[2]["age"] = "22"
		usinfo[2]["gender"] = "男"
	}
	fmt.Println(usinfo)

	for _, v := range usinfo {
		for key, value := range v {
			fmt.Printf("key:%v value:%v", key, value)
		}
		fmt.Println()
	}

	// asinfo := make(map[string]string)
	// asinfo["name"] = "zs"
	// asinfo["age"] = "18"
	// asinfo["gender"] = "男"

	asinfo := make(map[string][]string)
	asinfo["hobby"] = []string{"篮球", "足球", "乒乓球"}
	asinfo["work"] = []string{"php", "java", "goland"}
	fmt.Println(asinfo)

	for _, v := range asinfo {
		for _, vv := range v {
			fmt.Println(vv)
		}
	}

	// map类型也是引用类型
	// 引用类型：改变变量副本的值，会影响到原始变量的值
	// 值类型：改变变量副本的值，不会影响到原始变量的值
	// map类型也是引用类型
	m1 := make(map[string]string)
	m1["name"] = "zs"
	m1["age"] = "18"
	m1["gender"] = "男"

	m2 := m1
	m2["name"] = "ls"
	m2["age"] = "20"
	m2["gender"] = "女"

	fmt.Println(m1)
	fmt.Println(m2)

	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)

	for key, value := range map1 {
		fmt.Println(key, value)
	}
	// 按照key升序输出
	// 1 13
	// 4 56
	// 8 90
	// 10 100
	var ketSlice []int
	for key := range map1 {
		ketSlice = append(ketSlice, key)
	}
	sort.Ints(ketSlice)
	fmt.Println(ketSlice)
	for _, v := range ketSlice {
		fmt.Println(v, map1[v])
	}

	// 写一个程序统计字符串中单词出现次数
	str := "how do you do i do you do do do do do how ww w z z z z c c c "
	// 1. 定义一个map，key为单词，value为单词出现的次数
	wordcount := strings.Split(str, " ")
	wordmap := make(map[string]int)
	for _, v := range wordcount {
		wordmap[v]++
	}
	fmt.Println(wordmap)
}

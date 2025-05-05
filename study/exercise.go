package main

import (
	"fmt"
	"sort"
)

// 使用 var 关键字声明全局变量
var a = "全局变量"

// sortIntAsc 对整型切片进行升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

// sortIntDesc 对整型切片进行降序排序
func sortIntDesc(slice []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice
}

func sortMapAsc(m map[string]string) map[string]string {
	// 对map的key进行排序
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 按照排序后的key重新构建map

	sortedMap := make(map[string]string)
	for _, k := range keys {
		sortedMap[k] = m[k]
	}
	return sortedMap
}

func run() {
	fmt.Println(a)
}

func main() {
	// 排序封装成方法，实现整型切片的升序降序排序排列
	slice := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}

	// 升序排序
	ascSlice := sortIntAsc(slice)
	fmt.Println("升序排序结果:", ascSlice)

	// 降序排序
	descSlice := sortIntDesc(slice)
	fmt.Println("降序排序结果:", descSlice)

	m1 := map[string]string{
		"name":    "zs",
		"age":     "18",
		"gender":  "男",
		"address": "北京",
		"phone":   "13800138000",
	}
	// 对map的key进行排序
	sortedMap := sortMapAsc(m1)
	fmt.Println("排序后的map:", sortedMap)

	fmt.Println(a)
	run()
}

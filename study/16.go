// 数组切片排序以及sort
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 选择排序
	// 选择排序是一种简单直观的排序算法，它的基本思想是：
	// 第一次从 arr[0]~arr[n-1]中选取最小值，与 arr[0]交换，
	// 第二次从 arr[1]~arr[n-1]中选取最小值，与 arr[1]交换，
	// 第三次从 arr[2]~arr[n-1]中选取最小值，与 arr[2]交换，
	// …，
	numSlice := []int{9, 3, 4, 67, 1}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] > numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}
		}
	}
	fmt.Println(numSlice)

	// 从大到小
	numSlice1 := []int{9, 3, 4, 67, 1}
	for i := 0; i < len(numSlice1); i++ {
		for j := i + 1; j < len(numSlice1); j++ {
			if numSlice1[i] < numSlice1[j] {
				temp := numSlice1[i]
				numSlice1[i] = numSlice1[j]
				numSlice1[j] = temp
			}
		}
	}
	fmt.Println(numSlice1)

	// 冒泡排序
	// 冒泡排序是一种简单的排序算法，它的基本思想是：
	// 第一次从 arr[0]~arr[n-1]中选取最大值，与 arr[n-1]交换，
	// 第二次从 arr[0]~arr[n-2]中选取最大值，与 arr[n-2]交换，
	// 第三次从 arr[0]~arr[n-3]中选取最大值，与 arr[n-3]交换，
	// …，

	numSlice2 := []int{9, 3, 4, 67, 1}
	for i := 0; i < len(numSlice2); i++ {
		for j := 0; j < len(numSlice2)-i-1; j++ {
			if numSlice2[j] > numSlice2[j+1] {
				temp := numSlice2[j]
				numSlice2[j] = numSlice2[j+1]
				numSlice2[j+1] = temp
			}
		}
	}
	fmt.Println(numSlice2)

	intList := []int{93213, 3321, 14, 637, 13}
	float8List := []float64{9.11, 3.23, 4.333, 11.3213, 67.2312, 1.021321}
	stringList := []string{"a", "s", "c", "g", "e"}
	sort.Ints(intList)
	sort.Float64s(float8List)
	sort.Strings(stringList)
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)

	// sort降序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	sort.Sort(sort.Reverse(sort.Float64Slice(float8List)))
	sort.Sort(sort.Reverse(sort.StringSlice(stringList)))
	fmt.Println(intList)
	fmt.Println(float8List)
	fmt.Println(stringList)
}

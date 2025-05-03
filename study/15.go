package main

import "fmt"

func main() {
	// 切片
	// 1. 切片是数组的引用
	// 2. 切片是一个引用类型
	// 3. 切片是一个可变长度的数组
	/*
		    值类型：改变变量副本的值，不会影响到原始变量的值
			引用类型：改变变量副本的值，会影响到原始变量的值
	*/
	var arr1 []int
	fmt.Printf("arr1=%v 类型:%T\n", arr1, arr1)
	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr2=%v 类型:%T\n", arr2, arr2)
	arr3 := []int{1: 2, 3: 4, 5: 6}
	fmt.Printf("arr3=%v 类型:%T\n", arr3, arr3)
	arr4 := make([]int, 5, 10)
	fmt.Printf("arr4=%v 类型:%T\n", arr4, arr4)
	// go中声明切片之后， 默认值为nil
	// nil 切片的长度和容量都是0
	fmt.Println(len(arr4))
	fmt.Println(cap(arr4))
	a := []string{"php", "java", "goland"}
	for i := 0; i < len(a); i++ {
		fmt.Printf("key:%v value:%v\n", i, a[i])
	}
	for k, v := range a {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	s := [5]int{1, 2, 3, 4, 5}
	b := s[:] // 获取数组里面的所有值
	fmt.Printf("b=%v 类型:%T\n", b, b)
	b[0] = 100
	fmt.Printf("b=%v 类型:%T\n", b, b)

	c := s[1:3] // 获取数组里面的一部分值
	fmt.Printf("c=%v 类型:%T\n", c, c)
	d := s[1:] // 获取数组里面的一部分值
	fmt.Printf("d=%v 类型:%T\n", d, d)
	e := s[:3] // 获取数组里面的一部分值
	fmt.Printf("e=%v 类型:%T\n", e, e)
	f := s[1:3:5] // [2 3]
	fmt.Printf("f=%v 类型:%T\n", f, f)

	// 基于切片创建切片
	as := []string{"php", "java", "goland", "nodejs", "rust"}
	bs := as[1:3]
	fmt.Printf("b=%v 类型:%T\n", bs, bs)

	// 切片的长度和容量
	fmt.Println(len(bs))
	fmt.Println(cap(bs))
	// 长度就是他所含元素个数
	// 容量就是从切片的第一个元素开始，到底层数组的最后一个元素的个数

	// make函数创建切片
	// make([]类型，长度，容量)
	// make([]类型，长度)
	// make([]类型)
	// make([]类型，容量)

	sliceA := make([]int, 4, 10)
	fmt.Printf("sliceA=%v 类型:%T\n", sliceA, sliceA)
	fmt.Println(len(sliceA), cap(sliceA))
	sliceA[0] = 100
	fmt.Printf("sliceA=%v 类型:%T\n", sliceA, sliceA)

	sliceB := []string{"php", "java", "goland", "nodejs", "rust"}
	sliceB[2] = "js"
	fmt.Printf("sliceB=%v 类型:%T\n", sliceB, sliceB)

	// go中不可以使用下标扩容只能使用append函数
	var sliceC []int
	fmt.Printf("长度:%v 容量:%v\n", len(sliceC), cap(sliceC))
	sliceC = append(sliceC, 10)
	fmt.Printf("sliceC=%v 类型:%T\n", sliceC, sliceC)
	sliceC = append(sliceC, 20)
	fmt.Printf("sliceC=%v 类型:%T\n", sliceC, sliceC)

	// append合并切片
	sliceD := []int{1, 2, 3, 4, 5}
	sliceE := []int{6, 7, 8, 9, 10}
	sliceF := append(sliceD, sliceE...)
	fmt.Printf("sliceF=%v 类型:%T 长度:%v 容量:%v\n", sliceF, sliceF, len(sliceF), cap(sliceF))

	// 切片扩容策略
	// 1. 如果原切片的容量小于1024，那么新切片的容量就是原切片的容量的2倍
	// 2. 如果原切片的容量大于等于1024，那么新切片的容量就是原切片的容量的1.25倍
	// 3. 如果原切片的容量大于等于1024，并且新切片的容量大于等于1024，那么新切片的容量就是原切片的容量的1.25倍
	// 4. 如果原切片的容量大于等于1024，并且新切片的容量小于1024，那么新切片的容量就是原切片的容量的1.25倍
	// 5. 如果原切片的容量小于1024，并且新切片的容量小于1024，那么新切片的容量就是原切片的容量的2倍

	var sliceSa []int
	for i := 1; i <= 10; i++ {
		sliceSa = append(sliceSa, i)
		fmt.Printf("sliceS=%v 类型:%T 长度:%v 容量:%v\n", sliceSa, sliceSa, len(sliceSa), cap(sliceSa))
	}

	sa := []int{1, 2, 3, 4, 5}
	sb := sa
	sb[0] = 100
	fmt.Printf("sa=%v 类型:%T\n", sa, sa)
	fmt.Printf("sb=%v 类型:%T\n", sb, sb)

	// copy函数
	slicea := []int{1, 2, 3, 4, 5}
	sliceb := make([]int, 5, 10)
	copy(sliceb, slicea)
	fmt.Printf("slicea=%v 类型:%T\n", slicea, slicea)
	fmt.Printf("sliceb=%v 类型:%T\n", sliceb, sliceb)
	sliceb[0] = 100
	fmt.Printf("slicea=%v 类型:%T\n", slicea, slicea)
	fmt.Printf("sliceb=%v 类型:%T\n", sliceb, sliceb)

	aaa := []int{1, 2, 3, 4, 5}
	// 删除索引为2的元素 删除的元素为3
	aaa = append(aaa[:2], aaa[3:]...)
	fmt.Printf("aaa=%v 类型:%T\n", aaa, aaa)
	// 删除4和5
	aaa = append(aaa[:2])
	fmt.Printf("aaa=%v 类型:%T\n", aaa, aaa)

	s1 := "big"
	bytestr := []byte(s1)
	fmt.Printf("bytestr=%v 类型:%T\n", bytestr, bytestr)
	bytestr[0] = 'p'
	fmt.Printf("bytestr=%v 类型:%T\n", bytestr, bytestr)

	s2 := "你好goland"
	runstr := []rune(s2)
	fmt.Printf("runtstr=%v 类型:%T\n", runstr, runstr)
	runstr[0] = '大'
	fmt.Printf("runtstr=%v 类型:%T\n", runstr, runstr)
	fmt.Println(string(runstr))
}

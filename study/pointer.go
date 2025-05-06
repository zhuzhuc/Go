package main

import "fmt"

// fn1 函数通过值传递接收一个整数，修改其参数值，但不会影响原始变量
func fn1(x int) {
	x = 10
}

// fn2 函数通过指针传递接收一个整数指针，修改指针指向的值
func fn2(x *int) {
	*x = 40
}

func main() {
	fmt.Println("pointer")
	var a int = 10
	fmt.Printf("a的值: %v a的类型:%T a的地址是%p\n", a, a, &a)

	p := &a
	fmt.Printf("p的值: %v p的类型:%T p的地址是%p\n", p, p, &p)
	fmt.Printf("p指向的值: %v p指向值的类型:%T p指向值的地址是%p\n", *p, *p, &*p) // *p表示取出p这个变量对应的内存地址的值

	*p = 20
	fmt.Printf("a的值: %v a的类型:%T a的地址是%p\n", a, a, &a)
	fmt.Printf("p的值: %v p的类型:%T p的地址是%p\n", p, p, &p)

	s := 5
	fn1(s)
	fmt.Println(s)

	fn2(&s)
	fmt.Println(s)

	userinfo := make(map[string]string)
	userinfo["name"] = "zhangsan"
	userinfo["age"] = "18"
	userinfo["sex"] = "male"
	fmt.Println(userinfo)

	as := make([]int, 5, 10)
	fmt.Println(as)
	fmt.Printf("as的长度: %v as的容量: %v\n", len(as), cap(as))
	as = append(as, 1, 2, 3, 4, 5)
	fmt.Println(as)

	// var a *int//指针也是引用数据类型
	// *a = 100
	// fmt.Println(a)

	aa := new(int)
	bb := new(bool)
	fmt.Printf("%T\n", aa)
	fmt.Printf("%T\n", bb)
	fmt.Println(*aa)
	fmt.Println(*bb)

	var ass *int
	ass = new(int)
	*ass = 100
	fmt.Println(*ass)

	f := new(bool)
	fmt.Println(*f)
}

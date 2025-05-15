package main

import (
	"fmt"
	"reflect"
)

// 定义自定义类型
type (
	// myINT 是 int 类型的别名
	myINT int
	// Person 是一个结构体类型，包含姓名和年龄字段
	Person struct {
		Name string
		Age  int
	}
)

// reflectType 使用反射获取任意变量的类型
func reflectType(x interface{}) {
	// 获取变量的类型
	// 1. 首先获取 reflect.Type 类型
	// 2. 然后通过 reflect.Type 类型调用 String() 方法，可获取类型的名称
	// 3. 也可以通过 reflect.Type 类型调用 Kind() 方法，获取类型的种类
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v ------kind: %v\n", t.Name(), t.Kind())
	// fmt.Println(t.String())
}

// reflectSetValue 使用反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 检查传入的是否为指针类型，如果不是则打印错误信息并返回
	if v.Kind() != reflect.Ptr {
		fmt.Println("错误：必须传入指针类型")
		return
	}
	// 获取指针指向的值
	elem := v.Elem()
	// 检查指针指向的值是否有效且可设置，如果不是则打印错误信息并返回
	if !elem.IsValid() || !elem.CanSet() {
		fmt.Println("错误：无效或不可寻址的值")
		return
	}

	// 根据指针指向的值的类型进行不同的设置操作
	switch elem.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// 设置整型值
		elem.SetInt(20)
	case reflect.String:
		// 设置字符串值
		elem.SetString("zzc")
	default:
		// 不支持的类型，打印类型信息
		fmt.Printf("不支持的类型: %v\n", elem.Kind())
	}
}

// reee 使用反射获取变量的原始值并判断其类型
func reee(x interface{}) {
	// bb, ok := x.(int)
	// num := 10 + bb
	// if ok {
	//  fmt.Println(num)
	// }
	// 反射实现
	// v := reflect.ValueOf(x)
	// fmt.Println(v)
	// 使用反射获取变量的原始值
	v := reflect.ValueOf(x)
	// 计算值并打印
	m := v.Int() + 12
	fmt.Println(m)

	// 获取变量的类型种类
	kind := v.Kind()
	// 根据类型种类进行判断并打印类型信息
	switch kind {
	case reflect.Int:
		fmt.Println("int")
	case reflect.Float32:
		fmt.Println("float32")
	case reflect.String:
		fmt.Println("string")
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.Bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	// 定义不同类型的变量
	a := 10
	b := 10.1
	c := "zzc"
	d := true
	e := []int{1, 2, 3}
	f := map[string]int{"zzc": 1, "zzc1": 2}
	g := struct {
		Name string
		Age  int
	}{"zzc", 18}
	h := &g

	// 调用 reflectType 函数测试不同类型的变量
	reflectType(a)
	reflectType(b)
	reflectType(c)
	reflectType(d)
	reflectType(e)
	reflectType(f)
	reflectType(g)
	reflectType(h)

	// 测试自定义类型
	var ok myINT = 34
	reflectType(ok)
	var ff Person = Person{"zzc", 18}
	reflectType(ff)

	// 测试 reee 函数
	aaaa := 13
	reee(aaaa)

	// 测试 reflectSetValue 函数设置 int64 类型的值
	var ass int64 = 10
	reflectSetValue(&ass)
	fmt.Println(ass)

	// 测试 reflectSetValue 函数设置 string 类型的值
	var aab string = "zzsasdadsac"
	reflectSetValue(&aab)
	fmt.Println(aab)
}

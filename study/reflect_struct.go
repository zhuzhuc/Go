package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float64 `json:"score"`
}

func Print() {
	fmt.Println("Print()")
}

// GetInfo 打印学生信息
func (s Student) GetInfo() {
	fmt.Println("Name = ", s.Name, "Age = ", s.Age, "Score = ", s.Score)
}

// SetInfo 设置学生信息
func (s *Student) SetInfo(name string, age int, score float64) {
	s.Name = name
	s.Age = age
	s.Score = score
}

// PrintInfo 打印学生信息
func (s Student) PrintInfo() {
	fmt.Println("Name = ", s.Name, "Age = ", s.Age, "Score = ", s.Score)
}

// PrintStruct 使用反射打印结构体的字段和方法信息
func PrintStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 如果传入的是指针，获取指针指向的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Println("错误：必须传入结构体类型")
		return
	}

	// 打印结构体的字段信息
	fmt.Println("结构体字段信息：")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("字段名: %s, 类型: %v, 值: %v, 标签: %v\n", field.Name, field.Type, value.Interface(), field.Tag.Get("json"))
	}

	// 打印结构体的方法信息
	fmt.Println("结构体方法信息：")
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("方法名: %s, 类型: %v\n", method.Name, method.Type)
	}

	// 调用 GetInfo 方法
	infoMethod := v.MethodByName("GetInfo")
	if infoMethod.IsValid() {
		infoMethod.Call(nil)
	} else {
		fmt.Println("未找到 GetInfo 方法")
	}

	// 若传入的是值类型，将其转换为指针类型，以便调用指针接收者方法
	if !v.CanAddr() {
		v = reflect.New(t).Elem()
		v.Set(reflect.ValueOf(s))
	}

	// 调用 SetInfo 方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf("zdaszc"))
	params = append(params, reflect.ValueOf(2320))
	params = append(params, reflect.ValueOf(3299.9))
	setInfoMethod := v.Addr().MethodByName("SetInfo")
	if setInfoMethod.IsValid() {
		setInfoMethod.Call(params)
	} else {
		fmt.Println("未找到 SetInfo 方法")
	}
}

func refectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("错误：必须传入指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("错误：必须传入结构体类型")
		return
	}
	// 修改结构体的字段值
	name := v.Elem().FieldByName("Name")
	name.SetString("zzc zzc")

	age := v.Elem().FieldByName("Age")
	age.SetInt(23330)
	score := v.Elem().FieldByName("Score")
	score.SetFloat(99323.9)
}

func main() {
	stu := Student{
		Name:  "zzc",
		Age:   20,
		Score: 99.9,
	}
	stu.GetInfo()
	stu.SetInfo("zsadszc", 202132, 932139.9)
	stu.PrintInfo()

	// 调用 PrintStruct 函数
	PrintStruct(stu)
	fmt.Println("----------------------------------------------------------------------")

	refectChangeStruct(&stu)
	fmt.Println("修改后的结构体值：", stu)
}

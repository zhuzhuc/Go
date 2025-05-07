package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` // 结构体标签，用于指定json序列化时的字段名
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

type Class struct {
	Title    string    `json:"title"`
	Students []Student `json:"students"`
}

func main() {
	s1 := Student{
		ID:     12,
		Gender: "男",
		Name:   "张三",
		Sno:    "s001",
	}
	fmt.Printf("%#v\n", s1)
	//结构体转json
	//jsonStr, err := json.Marshal(s1)
	//if err != nil {
	//	fmt.Println("json marshal failed")
	//	return
	//}
	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("jsonStr=%v\n", jsonStr)

	// json转结构体
	var s2 Student
	err := json.Unmarshal(jsonByte, &s2)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("%#v\n", s2)

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	c := Class{
		Title:    "1班",
		Students: make([]Student, 0),
	}
	for i := 1; i <= 10; i++ {
		s := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("学生%d", i),
		}
		c.Students = append(c.Students, s) // 将s添加到c.Students切片中
	}
	fmt.Println(c)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	strByte, _ := json.Marshal(c)
	str := string(strByte)
	fmt.Println(str)

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	str1 := `{"title":"1班","students":[{"id":1,"gender":"男","name":"学生1","sno":""},{"id":2,"gender":"男","name":"学生2","sno":""},{"id":3,"gender":"男","name":"学生3","sno":""},{"id":4,"gender":"男","name":"学生4","sno":""},{"id":5,"gender":"男","name":"学生5","sno":""},{"id":6,"gender":"男","name":"学生6","sno":""},{"id":7,"gender":"男","name":"学生7","sno":""},{"id":8,"gender":"男","name":"学生8","sno":""},{"id":9,"gender":"男","name":"学生9","sno":""},{"id":10,"gender":"男","name":"学生10","sno":""}]}`
	ss := &Class{}
	err = json.Unmarshal([]byte(str1), ss)
	if err != nil {
		fmt.Println("json unmarshal failed")
	} else {
		fmt.Printf("%#v\n", ss)
		fmt.Printf("ss.Title=%v\n", ss.Title)
	}
}

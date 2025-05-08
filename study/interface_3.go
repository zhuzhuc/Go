package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// goland中空接口和类型断言使用细节
func main() {
	userinfo := make(map[string]interface{})

	userinfo["name"] = "tom"
	userinfo["age"] = 18
	// 保存原始的地址切片
	originalAddress := []string{"北京", "上海", "深圳"}
	userinfo["address"] = originalAddress

	fmt.Println(userinfo)

	address := Address{
		Name:  "tom",
		Phone: 123456,
	}
	// 使用不同的键存储 Address 类型的数据
	userinfo["addressInfo"] = address
	fmt.Println(address.Name)

	// 如何取出map中的值
	// 1.先将map转成interface{}
	var userinfo2 interface{} = userinfo
	// 2.将interface{}转成需要的类型
	// 2.1先将interface{}转成map
	userinfo3, ok := userinfo2.(map[string]interface{})
	if !ok {
		fmt.Println("类型断言失败，不是 map[string]interface{} 类型")
		return
	}
	// 2.2取出对应的值
	name, ok := userinfo3["name"].(string)
	if !ok {
		fmt.Println("类型断言失败，name 不是 string 类型")
		return
	}
	age, ok := userinfo3["age"].(int)
	if !ok {
		fmt.Println("类型断言失败，age 不是 int 类型")
		return
	}
	// 取出 Address 类型的数据
	addressInfo, ok := userinfo3["addressInfo"].(Address)
	if !ok {
		fmt.Println("类型断言失败，addressInfo 不是 Address 类型")
		return
	}

	fmt.Println(name, age, addressInfo)

	// 取出原始的地址切片
	hobby3, ok := userinfo3["address"].([]string)
	if !ok {
		fmt.Println("类型断言失败，address 不是 []string 类型")
		return
	}
	fmt.Println(hobby3[0])

	address2, _ := userinfo3["addressInfo"].(Address)
	fmt.Println(address2.Phone)
}

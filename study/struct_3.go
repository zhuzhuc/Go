package main

import "fmt"

type Person struct {
	string
	int
}

type PPerson struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

// 结构体嵌套
type User struct {
	Username string
	Password string
	Address  // Address // 表示user结构体中嵌套了一个Address结构体
	Email    Email
}

// 依旧可以匿名嵌套
type Address struct {
	Name    string
	Phone   string
	City    string
	Addtime string
}

// 当访问结构体成员时会先在结构体中查找该字段， 找不到再去匿名结构体中查找。
type Email struct {
	Account string
	Addtime string
}

// 结构体的继承
// 父结构体
type Anmial struct {
	Name string
}

// run 方法，打印动物在跑的信息
func (a Anmial) run() {
	fmt.Println("动物在跑", a.Name)
}

type Dog struct {
	Age    int
	Anmial // 表示dog结构体中嵌套了一个Anmial结构体，dog结构体继承了Anmial结构体的字段和方法
}

// 也可以嵌套指针
// wang 方法，打印小狗在叫的信息
func (d Dog) wang() {
	fmt.Println("小狗在叫", d.Name)
}

func main() {
	p1 := Person{
		"哈哈",
		18,
	}
	fmt.Println(p1)

	var p2 PPerson
	p2.Name = "哈哈"
	p2.Age = 18
	p2.Hobby = []string{"java", "php", "go"}
	p2.map1 = make(map[string]string)
	p2.map1["name"] = "zzc"
	p2.map1["phone"] = "value2"
	p2.map1["key3"] = "value3"
	fmt.Printf("%#v\n", p2)

	fmt.Println("---------------------------------------------------------------------------------------")
	var u User
	u.Username = "zzc"
	u.Password = "123456"
	u.Address.Name = "嘻嘻"
	u.Address.Phone = "123456"
	u.Address.City = "beijing"
	u.City = "shanghai"
	// u.Addtime = "2024-07-23 15:04:05"
	u.Address.Addtime = "2099-07-23 15:04:05"
	u.Email.Addtime = "2024-07-23 15:04:05"
	fmt.Println(u)
	fmt.Println(u.Address.Phone)
	fmt.Println("---------------------------------------------------------------------------------------")

	// 修正结构体初始化语法和类型错误
	d := Dog{
		Age: 1,
		Anmial: Anmial{
			Name: "saddog",
		},
	}
	d.wang()
	d.run()

	fmt.Println("---------------------------------------------------------------------------------------")
}

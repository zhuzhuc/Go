package main

import "fmt"

// Usber 定义一个接口，包含 Start 和 Stop 方法
type Usber interface {
	Start()
	Stop()
}

// Computer 定义计算机结构体
type Computer struct{}

// work 方法接收一个实现了 Usber 接口的参数，调用其 Start 和 Stop 方法
func (c Computer) work(usb Usber) {
	usb.Start()
	usb.Stop()
}

// Phone 定义手机结构体
type Phone struct {
	Name string
}

// Start 手机开始工作
func (p Phone) Start() {
	fmt.Println(p.Name, "start")
}

// Stop 手机停止工作
func (p Phone) Stop() {
	fmt.Println(p.Name, "stop")
}

// Camera 定义相机结构体
type Camera struct {
	// Name string
}

// Start 相机开始工作
func (c Camera) Start() {
	fmt.Println("camera start")
}

// Stop 相机停止工作
func (c Camera) Stop() {
	fmt.Println("camera stop")
}

func main() {
	computer := Computer{}
	phone := Phone{Name: "xiaomi"}
	camera := Camera{}
	computer.work(phone)
	computer.work(camera)
}

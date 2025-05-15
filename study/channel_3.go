package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello,world", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "goland" // error
}

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20

	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)

	// ch2 := make(chan<- int, 2)
	// ch2 <- 10
	// ch2 <- 20
	// // 不能从ch2中读取数据，因为ch2是一个只能写的channel
	// // m3 := <-ch2
	// // m4 := <-ch2
	// // fmt.Println(m3, m4)
	// ch3 := make(<-chan int, 2)
	// // 不能向ch3中写入数据，因为ch3是一个只能读的channel
	// // ch3 <- 10
	// // ch3 <- 20
	// m5 := <-ch3
	// m6 := <-ch3
	// fmt.Println(m5, m6)

	// 多路复用
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 使用sellect不需要关闭channel，因为select会自动判断channel是否关闭
	for {
		select {
		case v := <-intChan:
			fmt.Println("从intChan中读取的数据", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Println("从stringChan中读取的数据", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕，退出")
			return
		}
	}

	go sayHello()
	go test()
	time.Sleep(time.Millisecond * 5000)
}

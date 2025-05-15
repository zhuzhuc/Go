package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
	m1    = make(map[int]int, 10)
)

func test(num int) {
	mutex.Lock()
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m1[num] = sum
	fmt.Printf("m1[%d] = %d\n", num, sum)
	time.Sleep(time.Millisecond * 50)
	mutex.Unlock()
	wg.Done()
}

// 读写互斥锁
// 读写互斥锁就是将读写操作的锁分成了两个锁，一个读操作的锁，一个写操作的锁
// 通过4个goroutine来演示读写互斥锁的用法
// 分别对全局变量进行读取和写入操作
// 在读取操作的函数中，加入了读锁
// 在写入操作的函数中，加入了写锁
// 在主函数中，分别启动40个goroutine进行读取操作，10个goroutine进行写入操作
// 在主函数中，等待所有的goroutine执行完毕
// 在主函数中，打印出全局变量的值
// 在主函数中，打印出主函数的结束信息
// 在主函数中，等待所有的goroutine执行完毕
// 在主函数中，打印出主函数的结束信息
func write() {
	mutex.Lock()
	fmt.Println("----write() 开始执行")

	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func read() {
	mutex.Lock()
	fmt.Println("-----read() 开始执行")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 40; r++ {
		wg.Add(1)
		go test(r)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println("main() 退出")
}

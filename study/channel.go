package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(sss chan int) {
	for i := 1; i <= 10; i++ {
		sss <- i
		fmt.Printf("fn1() 写入的第%v条数据成功\n", i)
		time.Sleep(time.Millisecond * 1000)
	}
	close(sss) // 关闭channel
	wg.Done()
}

func fn2(sss chan int) {
	for v := range sss {
		fmt.Printf("fn2() 读取的第%v条数据成功\n", v)
	}
	fmt.Println()
	wg.Done()
}

func main() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan []int
	var ch4 chan map[int]string
	var ch5 chan *int
	var ch6 chan<- int
	var ch7 <-chan int
	fmt.Println(ch1, ch2, ch3, ch4, ch5, ch6, ch7)

	ch := make(chan int, 3)       // 带缓冲的channel，缓冲区大小为3
	fmt.Println(len(ch), cap(ch)) // 0 3
	ch <- 10
	fmt.Println(len(ch), cap(ch)) // 1 3
	ch <- 20
	ch <- 30
	fmt.Println(len(ch), cap(ch)) // 2 3
	m := <-ch
	fmt.Println(m)

	<-ch // 从channel中取出一个元素，但是不赋值给变量
	d := <-ch
	fmt.Println(d)

	mm := make(chan int, 4)
	mm <- 14
	mm <- 2432
	mm <- 33

	ch22 := mm
	ch22 <- 10
	<-mm
	<-mm
	<-mm
	ok := <-mm
	fmt.Println(ok)
	// 管道是引用类型，必须初始化才能写入数据，即make后才能使用。
	// 管道是有类型的，一个string的channel只能存放string类型的数据。
	// channel是引用类型，在make的时候，必须指定大小。

	// 管道的遍历
	ch33 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch33 <- i
	}
	close(ch33) // 关闭channel
	for v := range ch33 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	// 管道阻塞
	// 当我们给一个channel中发送数据的时候，在没有接收方接收数据的时候，会发生阻塞。
	// 当我们从一个channel中接收数据的时候，在没有发送方发送数据的时候，也会发生阻塞。
	// 当我们关闭一个channel的时候，会发生阻塞。
	// 当我们从一个关闭的channel中接收数据的时候，会发生阻塞。
	// ch6 := make(chan int 1)
	// ch6 <- 10
	// ch6 <- 20 // all goroutines are asleep - deadlock!

	ch77 := make(chan string, 2)
	ch77 <- "hello"
	ch77 <- "world"

	m1 := <-ch77
	m2 := <-ch77
	// m3 := <-ch77 // all goroutines are asleep - deadlock!
	fmt.Println(m1)

	fmt.Println(m2)

	// for循环遍历的时候管道可以不关闭
	chh := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		chh <- i
	}
	for j := 1; j <= 10; j++ {
		fmt.Printf("%v ", <-chh)
	}
	fmt.Println()

	sss := make(chan int, 10)
	wg.Add(1)
	go fn1(sss)
	go fn2(sss)
	wg.Wait()
	fmt.Println("main() 退出")
}

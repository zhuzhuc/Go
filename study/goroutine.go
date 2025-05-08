package main

// 在主线程中，开启一个goroutine，该协程每隔1秒输出“hello,world”
// 在主线程中也每隔一秒输出“hello,golang”，输出10次后，退出程序
// 要求主线程和goroutine同时执行
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 注意：主线程执行完毕后即使协程没有执行完毕程序也会退出
// 协程可以在主线程没有执行完毕前提前退出，协程是否执行完毕不会影响主线程的执行
// 为了保证我们的程序可以顺利执行。我们想让协程执行完毕后在执行主进程退出，这个时候我们可以使用sync.WaitGroup来实现

func test() {
	defer wg.Done() // 计数器-1
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello goland")
		time.Sleep(time.Millisecond * 50)
	}
}

func test1() {
	defer wg.Done() // 计数器-1
	for i := 0; i < 10; i++ {
		fmt.Println("test1() hello goland")
		time.Sleep(time.Millisecond * 500)
	}
}

var mm sync.WaitGroup

func test3(num int) {
	defer mm.Done() // 计数器-1
	for i := 0; i < 10; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	wg.Add(1)  // 计数器+2，因为开启了两个 goroutine
	go test()  // 开启一个协程，去执行test()函数，不要等待test()函数执行完毕，直接执行下面的代码
	wg.Add(1)  // 计数器+1，因为开启了一个 goroutine
	go test1() // 开启一个协程，去执行test1()函数，不要等待test1()函数执行完毕，直接执行下面的代码
	for i := 0; i < 10; i++ {
		fmt.Println("main() nihao goland")
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() // 等待计数器为0
	fmt.Println("main() 退出")

	// goland中设置CPU核数
	// 在goland中设置CPU核数，需要在goland的设置中找到CPU和内存，然后在CPU中设置CPU核数，设置完后，需要重启goland才能生效
	// 获取当前cpu的核数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)
	// 设置cpu的核数
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("OK")

	for i := 1; i <= 5; i++ {
		mm.Add(1)
		go test3(i)
	}
	mm.Wait()
	fmt.Println("main() 退出")
}

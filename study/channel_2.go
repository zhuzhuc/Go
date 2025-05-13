package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intchan chan int) {
	for i := 2; i < 122222220; i++ {
		intchan <- i
	}
	close(intchan)
	wg.Done()
}

func primeNum(intchan chan int, primechan chan int, exitChan chan bool) {
	for num := range intchan {
		var flag bool = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
		}
	}
	// 当intchan关闭后，primechan也需要关闭，否则会出现死锁
	exitChan <- true
	wg.Done()
}

func printPrime(primechan chan int) {
	// for prime := range primechan {
	// 	time.Sleep(time.Millisecond * 1000)
	// 	fmt.Println(prime)
	// }
	wg.Done()
}

func main() {
	start := time.Now().Unix()
	intchan := make(chan int, 1000)
	primechan := make(chan int, 1000000)
	exitChan := make(chan bool, 16)
	wg.Add(1)
	go putNum(intchan)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intchan, primechan, exitChan)
	}
	wg.Add(1)
	go printPrime(primechan)

	// 等待所有 16 个 primeNum 协程完成后，再关闭 primechan
	wg.Add(1) // 为匿名协程添加计数
	go func() {
		defer wg.Done() // 协程结束时减少计数
		for i := 0; i < 16; i++ {
			<-exitChan // 等待所有 16 个 primeNum 协程发送完成信号
		}
		close(primechan) // 所有协程完成后再关闭 primechan
	}()

	wg.Wait()

	end := time.Now().Unix()

	fmt.Println("main() 退出", end-start, "ms")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	timeobj := time.Now()
	fmt.Printf("timeobj = %v\n", timeobj)

	year := timeobj.Year()
	month := timeobj.Month()
	day := timeobj.Day()
	hour := timeobj.Hour()
	minute := timeobj.Minute()
	second := timeobj.Second()
	// 格式化时间
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	str := timeobj.Format("2006-01-02 03:04:05")
	// 解析时间
	fmt.Println(str)

	str2 := timeobj.Format("2006-01-02 15:04:05.000")
	fmt.Println(str2)

	// 时间戳
	timestamp := timeobj.Unix()
	fmt.Println("当前时间戳:", timestamp) // 这里实际是秒，不是毫秒

	unixNatime := timeobj.UnixNano() // ns
	fmt.Println("当前纳秒时间戳:", unixNatime)

	// 修改变量名避免重复声明
	str3 := "2024-07-23 15:04:05"
	// 使用解析结果
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", str3, time.Local)
	if err != nil {
		fmt.Println("解析时间出错:", err)
	} else {
		fmt.Println("解析后的时间戳:", parsedTime.Unix())
	}

	// 时间转时间戳
	timeobj2 := time.Now()
	// 使用赋值语句避免重复声明
	timestamp = timeobj2.Unix()
	fmt.Printf("timestamp = %v\n", timestamp)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Second)
	// 时间戳转时间
	// 使用赋值语句避免重复声明
	timeobj2 = time.Unix(timestamp, 0)
	fmt.Printf("timeobj2 = %v\n", timeobj2)
	// 时间间隔
	duration := timeobj2.Sub(timeobj)
	fmt.Println(duration)
	// 时间比较
	fmt.Println(timeobj.Before(timeobj2))
	fmt.Println(timeobj.After(timeobj2))
	fmt.Println(timeobj.Equal(timeobj2))
	// 时间加减
	timeobj3 := timeobj.Add(time.Hour)
	fmt.Printf("timeobj3 = %v\n", timeobj3)
	timeobj4 := timeobj.Add(-time.Hour)
	fmt.Printf("timeobj4 = %v\n", timeobj4)
	// 定时器
	ticker := time.NewTicker(time.Second)
	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
	}
	// 停止定时器
	ticker.Stop()

	fmt.Println("程序结束")
	time.Sleep(time.Second * 5)
	fmt.Println("程序结束")
	fmt.Println("程序结束")
	fmt.Println("程序结束")
	time.Sleep(time.Second * 5)
	fmt.Println("程序结束")
	fmt.Println("程序结束")
}

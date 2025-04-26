package main

import (
	"fmt"
)

func main() {
	var s string = ` 
	_____________________________  
	\____    /\____    /\_   ___ \ 
	  /     /   /     / /    \  \/ 
	 /     /_  /     /_ \     \____
	/_______ \/_______ \ \______  /	
	`
	fmt.Println(s)
	// 1 打印0-50之间的奇数
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	// 2 求1+2+3+...+100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 3. 打印1～100之间所有是9的倍数的整数的个数及总和
	count := 0
	sum2 := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum2 += i
		}
	}
	fmt.Println(count, sum2)

	// 计算5的阶乘
	sum3 := 1
	for i := 1; i <= 5; i++ {
		sum3 *= i
	}
	fmt.Println(sum3)

	// 打印一个n*m的矩形，
	var n, m int
	fmt.Println("please enter the n and m:")
	fmt.Scanln(&n, &m)
	num := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(num, " ")
			num++
		}
		fmt.Println()
	}

	// 打印一个杨辉三角
	var n2 int
	fmt.Println("please enter the n:")
	fmt.Scanln(&n2)

	// 创建一个二维切片来存储杨辉三角的值
	triangle := make([][]int, n2)
	for i := 0; i < n2; i++ {
		triangle[i] = make([]int, i+1)
		// 每行的第一个和最后一个元素为 1
		triangle[i][0] = 1
		triangle[i][i] = 1
		// 计算中间元素的值
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	// 打印杨辉三角
	for i := 0; i < n2; i++ {
		// 打印前导空格，使杨辉三角居中显示
		for k := 0; k < n2-i; k++ {
			fmt.Print("  ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%4d", triangle[i][j])
		}
		fmt.Println()
	}
}

/*
let mut n = String::new();
let mut m = String::new();
println!("please enter the n and m:");
io::stdin().read_line(&mut n).expect("Failed to read line");
io::stdin().read_line(&mut m).expect("Failed to read line");
let n: u32 = n.trim().parse().expect("Please type a number!");
let m: u32 = m.trim().parse().expect("Please type a number!");
let mut num = 1;
for _ in 1..=n {
    for _ in 1..=m {
        print!("{} ", num);
        num += 1;
    }
    println!();
}
*/

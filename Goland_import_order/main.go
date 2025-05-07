package main

import "goland_import_order/tools"

func init() {
	println("main init")
}

func main() {
	tools.PrintInfo()
}

// 在运行时， 被最后导入的包会最先初始化，并调用函数

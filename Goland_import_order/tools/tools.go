package tools

import "goland_import_order/calc"

func init() {
	println("tools init")
}

func PrintInfo() {
	calc.Add(10, 20)
	println("tools PrintInfo")
}

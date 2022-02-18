package main

import (
	"fmt"
)

/**
|--------------------------------------------------
| 高阶函数 1 函数作为参数 2 函数作为返回值
|--------------------------------------------------
*/
func sayhello(name string) {
	fmt.Println("...name", name)
}
func f1(name string, f func(string)) {
	f(name)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	f1("tom", sayhello)

	//函数做为返回值
	add := cal("+")
	r := add(1, 2)
	fmt.Printf("r: %v\n", r)

	fmt.Println("-----------")

	sub := cal("-")
	r = sub(100, 50)
	fmt.Printf("r: %v\n", r)
}

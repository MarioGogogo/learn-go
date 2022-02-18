package main

import "fmt"

/**
|--------------------------------------------------
| 函数类型 与 函数变量
|--------------------------------------------------
*/

func sum(a int, b int) int {
	return a + b

}

func max(a int, b int) {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	//函数类型
	type f1 func(int, int) int
	//相同函数类型  都可以赋值
	var ff f1
	ff = sum
	r := ff(1, 2)
	fmt.Println("...r", r)
}

package main

import "fmt"

/**
|--------------------------------------------------
| 递归
|--------------------------------------------------
*/
//斐波那契数列
func main() {
	r := f(5)
	fmt.Println("....r", r)
}

func f(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return f(n-1) + f(n-2)
}

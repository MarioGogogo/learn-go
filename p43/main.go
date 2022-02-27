package main

import "fmt"

/**
|--------------------------------------------------
| 匿名函数
|--------------------------------------------------
*/

func main() {
	//匿名函数
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}

	}
	i := max(1, 2)
	fmt.Println("....i", i)
	test()
}

//自己执行
func test() {
	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Println("...max", max)
	}(1, 2)
}

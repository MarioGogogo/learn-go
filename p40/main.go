package main

import "fmt"

/**
|--------------------------------------------------
|  golang 函数参数 必须指定参数类型
|--------------------------------------------------
*/
func f1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
func main() {
	r := f1(1, 2)
	fmt.Println("....r", r)

}

func f3(args ...int) {
	fmt.Println("...args", args)
}

/**
|--------------------------------------------------
|  map slice interface channel 修改会改变底层数据结构 会影响原数组
|--------------------------------------------------
*/

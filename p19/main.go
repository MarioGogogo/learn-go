package main

import "fmt"

/**
|--------------------------------------------------
|  流程控制
|--------------------------------------------------
*/

func main() {
	a := 100
	if a > 20 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}
	//下面这样的表达式  作用域 a不在外面 就是再if 语句内
	if a := 100; a > 20 {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}

	for i := 0; i < 4; i++ {
		fmt.Println("循环---", i)
	}

	x := [...]int{10, 20, 30}
	for i, v := range x {
		fmt.Println("v", v, i)
	}

}

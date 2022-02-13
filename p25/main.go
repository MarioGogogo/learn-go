package main

import "fmt"

/**
|--------------------------------------------------
|  for 循环  while 循环
|--------------------------------------------------
*/

func main() {
	f()
}

func f() {
	for i := 0; i < 10; i++ {
		fmt.Println("循环--", i)

	}

	j := 1
	for ; j < 10; j++ {
		fmt.Println("循环--", j)
	}

	//类似 js 中的 while
	x := 1
	for x <= 10 {
		fmt.Println("x----", x)
		x++ //结束条件
	}
}

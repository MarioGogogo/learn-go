package main

import "fmt"

/**
|--------------------------------------------------
| continue 循环中继续
|--------------------------------------------------
*/

func main() {
	//for 中
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("i", i)
		} else {
			continue
		}
	}
}

package main

import "fmt"

/**
|--------------------------------------------------
|    break 可以  for switch select
|--------------------------------------------------
*/

func main() {
	//break 跳出 for 循环
	for i := 0; i < 10; i++ {

		fmt.Println("i", i)
		if i >= 5 {
			break
		}
	}
	x := 2
	switch x {
	case 1:

		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		fallthrough //穿透
	case 3:
		fmt.Println("3")
		break
	}
}

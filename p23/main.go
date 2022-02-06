package main

import "fmt"

/**
|--------------------------------------------------
|  if嵌套
|--------------------------------------------------
*/

func main() {

	a := 100
	b := 200
	c := 3
	//多层嵌套
	if a > b {
		if a > c {
			fmt.Println("a 最大")
		}
	} else {
		if b > c {
			fmt.Println("b最大")
		} else {
			fmt.Println("c最大")
		}
	}

	//并集嵌套
	if a > b && a > c {
		fmt.Println("a 最大")
	} else if b > c {
		fmt.Println("b最大")
	} else {
		fmt.Println("c最大")
	}

}

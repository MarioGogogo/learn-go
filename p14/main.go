package main

import (
	"fmt"
)

func main() {
	var b1 bool = true
	b3 := true
	fmt.Println("b1", b1)
	fmt.Println("b3", b3)

	age := 15
	if age >= 17 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还是未成年")
	}

}

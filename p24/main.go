package main

import (
	"fmt"
)

func main() {
	f()
}

func f() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	case "C":
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
func test() {
	fmt.Println("hello vim")
}

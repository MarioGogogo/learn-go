package main

import "fmt"

/**
|--------------------------------------------------
| 运算符
|--------------------------------------------------
*/

func main() {
	a := 100
	b := 20

	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)

	fmt.Println("a*b=", a*b)

	fmt.Println("a/b=", a/b)
	fmt.Println("a%b=", a%b)

	a++
	fmt.Println("a++ ==", a)

	b--
	fmt.Println("b--", b)

}

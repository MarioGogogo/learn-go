package main

import "fmt"

/**
|--------------------------------------------------
| 函数 一等公民
|--------------------------------------------------
*/
func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}
func main() {
	r := sum(1, 2)
	fmt.Println("...r", r)
}

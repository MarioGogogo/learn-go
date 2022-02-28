package main

import "fmt"

/**
|--------------------------------------------------
| 指针
  go  记住2个符号 & 取地址  和 * 根据地址取值
|--------------------------------------------------
*/

func main() {
	var ip *int
	fmt.Println("...ip", ip)

	var i int = 100
	ip = &i
	fmt.Println("...ip", ip)
	fmt.Println("...p", *ip)
}

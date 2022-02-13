//数据类型  布尔  数字  字符串   派生(指针 数组 结构 切片)

package main

import "fmt"

func main() {
	var name string = "tom"
	age := 20
	b := true
	//打印类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	//指针类型
	a := 100
	p := &a
	fmt.Printf("%T\n", p)

	//数组
	list := [3]int{1, 2, 3}
	fmt.Printf("%T\n", list)

	//切片类型
	list1 := []int{1, 2, 3}
	fmt.Printf("%T\n", list1) //函数类型

}

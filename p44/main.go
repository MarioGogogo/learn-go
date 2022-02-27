package main

import "fmt"

/**
|--------------------------------------------------
| 闭包 闭包是将函数内部与外部 链接的桥梁
|--------------------------------------------------
*/
func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}

}
func main() {
	var f = add()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))
	fmt.Println(".).............")
	f1 := add()
	fmt.Println(f1(40))
	fmt.Println(f1(50))

	test1()
}

/**
|--------------------------------------------------
| 在f函数的生命周期内 变量x一直有效 (这句很重要)

 进阶
|--------------------------------------------------
*/
/**
  定义2个函数 返回值
*/
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(y int) int {
		base -= 1
		return base
	}
	return add, sub
}

func test1() {
	f1, f2 := calc(10)

	fmt.Println("....1", f1(1), f2(2))
	fmt.Println("....2", f1(3), f2(4))

}

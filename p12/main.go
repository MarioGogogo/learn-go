package main

import "fmt"

//常量
func main() {
	const constName string = "jack"

	const PI float64 = 3.14
	const PI2 = 3.1415 //可以省略类型

	const (
		ax = 100
		bx = 200
	)

	const i, j = 1, 2
	const a, b, c = 1, 2, "foo"

	//ioat 常量初始 0  i++
	const (
		x = iota
		y = iota
		z = iota
	)
	fmt.Println(x, y, z)

	//下滑线跳过
	const (
		x1 = iota
		_
		x2 = iota
	)
	fmt.Println(x1, x2)

	//中间跳过
	//下滑线跳过
	const (
		y1 = iota
		y2 = 100
		y3 = iota
	)
	fmt.Println(y1, y2, y3)
}

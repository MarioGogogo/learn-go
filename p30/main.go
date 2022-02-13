package main

import "fmt"

/**
|--------------------------------------------------
| 数组  长度不可修改
|--------------------------------------------------
*/

func main() {
	test()
	test2()
}
func test() {
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1:%T", a1)
	fmt.Printf("....a2:%T", a2)
	fmt.Println("a1,a2", a1, a2) //初始化0 值

}

//初始化
func test2() {
	var a = [3]int{1, 2}
	var b = [2]string{"tom"}
	var c = [3]bool{true, false}
	var d = [...]int{1, 3, 5} // 不用写具体长度
	var e = [...]int{0: 1, 3: 100, 5: 10}
	fmt.Printf("e:%v\n", e)
	fmt.Println('a', a)
	fmt.Println('b', b) //省略 默认 空
	fmt.Println('c', c) //省略值默认 false
	fmt.Println("d", d)
}

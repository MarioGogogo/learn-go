package main

import "fmt"

/**
|--------------------------------------------------
|  切片 可变长的数组  其他底层就是使用数组实现的 增加自动扩容功能
    切片拥有 相同类型元素的 可变长度序列
|--------------------------------------------------
*/

func main() {
	//初始化切片
	var names []string
	var numbers []int
	fmt.Println("name", names)
	fmt.Println("numbers", numbers)

	fmt.Println(names == nil)
	fmt.Println(numbers == nil)

}

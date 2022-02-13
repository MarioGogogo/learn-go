package main

import "fmt"

/**
|--------------------------------------------------
| range 循环

  数组 切片 字符串
   map 返回键和值
   通道 channel 值返回通道内的值
|--------------------------------------------------
*/

func f() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println("i", i, v)
	}
}

//切片
func f2() {
	var a = []int{1, 2, 3}
	for _, v := range a {
		fmt.Println("v():", v)
	}
}

//map
func f3() {
	m := make(map[string]string, 0)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "my@163.com"

	for key, value := range m {
		fmt.Printf("%v:%v,,", key, value)
	}

}
func main()
f()
f2()
f3()
}

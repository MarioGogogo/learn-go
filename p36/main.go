package main

import "fmt"

/**
|--------------------------------------------------
| map  里面key value 底层就是哈希表  无序的
|--------------------------------------------------
*/

func main() {
	//默认 map 是 nil
	//var map_variable map[key_data_type]value_data_type

	//使用make函数
	//map_variable = make(map[key_data_type]value_data_type)
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"
	fmt.Println("...m1.Name),", m1)

	m2 := map[string]string{
		"name":  "kite",
		"age":   "20",
		"email": "kite@gmail.com", //注意最后 又逗号
	}

	fmt.Println("...m2", m2) //无序的

	//访问
	fmt.Println("访问map", m2["age"])
}

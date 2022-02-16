package main

import "fmt"

/**
|--------------------------------------------------
| map 遍历
|--------------------------------------------------
*/

func main() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	for k := range m1 {
		fmt.Println("...k", k)
	}

	for k, v := range m1 {
		fmt.Println("...k,v", k, v)
	}

	for _, v := range m1 {
		fmt.Println("...k,v", v)
	}
}

package main

import "fmt"

/**
|--------------------------------------------------
|  切片循环
|--------------------------------------------------
*/

func main() {
	s1 := []int{1, 2, 3, 4}
	for i := 0; i < len(s1); i++ {
		fmt.Println("i,s1[i", i, s1[i])
	}
	for i, i2 := range s1 {
		fmt.Println("i,i2", i, i2)
	}
}

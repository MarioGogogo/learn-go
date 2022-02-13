package main

import "fmt"

/**
|--------------------------------------------------
| 数组 增删改查
|--------------------------------------------------
*/

func main() {
	s1 := []int{}
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5)

	fmt.Println("s1", s1)

	s3 := []int{3, 4, 5}
	s4 := []int{1, 2, 3}
	s4 = append(s3, s4...)
	fmt.Println("s4", s4)

	//del 删除的逻辑是这样的:  1 2 前 区间   写出要删除的索引3  链接索引+1 后面的元素 组成新数组
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println("s1", s1)
	//update
	s1[1] = 100
	fmt.Println("s1", s1)
	// query
	k := 100
	for i := 0; i < len(s1); i++ {
		if k == s1[i] {
			fmt.Println("找到了")
			break
		}
	}
}

package main

import "fmt"

// 佚名变量
func getNameAge() (string, int) {
	return "jack", 30
}
func main() {
	// var name string
	// var age int
	// var m bool
	//类型推断
	var name = "tom"
	var age = 10
	var m = true

	//批量初始化
	//var name,age,b = "tom",10,true

	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("m", m)

	jName, _ := getNameAge()
	fmt.Println("j_name", jName)

}

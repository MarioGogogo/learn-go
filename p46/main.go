package main

import "fmt"

/**
|--------------------------------------------------
| defer用途   会让后面跟随的语句延迟处理
 延迟处理的语句按defer定义的逆序执行  (js中先进后出 栈)
  关闭文件句柄
  锁资源释放
  数据库链接释放
|--------------------------------------------------
*/

func main() {
	fmt.Println("start")
	defer fmt.Println("setp1")
	defer fmt.Println("setp2")
	defer fmt.Println("setp3")
	fmt.Println("end")
}

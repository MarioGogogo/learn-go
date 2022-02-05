package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var str1 string = "hello go"
	var html string = `<html >
	< head > < title > hello
	golang </title > </head >
	</html >`

	fmt.Println("str1", str1)
	fmt.Println("str2", html)

	s1 := "hello"
	s2 := "golang2"

	fmt.Println(s1 + "," + s2)

	msg := strings.Join([]string{s1, s2}, ",")
	fmt.Println("msg", msg)

	//效率比较高的一种 buffer
	buf := bytes.Buffer{}
	buf.Write([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})

	fmt.Println("buffer---", buf.String())

	//转义

	s := "hello"

	fmt.Println(s + "\n" + "golang" + "\t")
	fmt.Println(s)

}

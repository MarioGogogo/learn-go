package main 

import "fmt"
import "strconv"
import "strings"
func main() {
	  var s1 string = "hello"
		var s2 string = "word"
		 i:= 10
		 pi:= &i

		 //常量
		 const (
			  one = iota+1
				two
		 )
		 str:= strings.Index("飞行高度","高")

	   i2s:=strconv.itoa(i)
		s2i,err:=strconv.Atoi(i2s)

    fmt.Println(s1,s2,":",s1+s2,i,"指针：",*pi,str,one,two,i2s,s2i,err)
}
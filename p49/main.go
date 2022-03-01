package main
import "fmt"
/**
|--------------------------------------------------
| 指向数组的指针
|--------------------------------------------------
*/

func main() {
	a :=[3]{1, 2, 4}
	var i int;
	var pa [3]*int;
	fmt.Println("...pa",pa);

	//赋值
	for i := 0; i < len(a); i++ {
	      pa[i] = &a[i]
		  fmt.Println("...pa",pa)
	}


}

package main 

import "fmt"

func main() {
    i:=6
		if i>10 {
			fmt.Println("i>10")
		}else if i>5&&i<=10{
			fmt.Println("5<i<=10")
		}else{
			fmt.Println("i<5")
		}

		sum:=0
		x:=1
		for{
			sum+=1
			x++
			if i>100 {
				break
			}
		}
		fmt.Println("sum",sum)
}




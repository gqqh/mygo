// control.go
package main

import (
	"fmt"
)

func main() {
	a := 10
/*
//	if a:= 1; a > 0{	//a := 1 初始化
	if a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println("a <= 0")
	}
*/
/*	for {	//while()
		a ++
		if a > 3 {
			break
		}
		fmt.Println(a)
	}*/
/*	for a >= 3 {	//while()
		a --
		fmt.Println(a)
	}
*/
/*	for i := 0; i < 3; i++ {	//for( ; ; )
		a++
		fmt.Println(a)
	}*/
	/*	
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("None")
	}
	*/
/*	
	switch a := 1; a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("None")
	}
	*/
/*	switch{
	case a < 1:
		fmt.Println("a < 1")
	case a > 10:
		fmt.Println("a > 10")
	default :
		fmt.Println("1 <= a <= 10")
	}
	*/

LABEL:
	for i := 0; i < 10; i++{
		for{
			fmt.Println(i)
			//continue LABEL
			break LABEL
			//goto LABEL
			
		}
	}
	
	fmt.Println(a)
}

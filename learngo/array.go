// array.go
package main

import (
	"fmt"
)

func main() {
	//	var a [2]int
	//	a := [20]int {19:1}
	/*	a := [...]int {19:1}
		var p *[20]int = &a
		var b [2]int
		x, y := 1, 2
		c := [...]*int {&x, &y}

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(p)
		fmt.Println(*p)
		fmt.Println(c)
	*/
	/*	a := [2]int{1, 2}
		b := [2]int{1, 3}
		fmt.Println(a == b)
	*/
	//buble sort
	a := [...]int{2, 3, 8, 7, 1, 0, 6, 11, 23, 43, 45, 99}
	fmt.Println(a)
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				tmp := a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
	fmt.Println(a)
}

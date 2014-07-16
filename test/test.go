package main

import "fmt"

func main(){
v1 := []int{2, 3}
	fmt.Println(v1)
	v1[0], v1[1] = v1[1], v1[0]
	fmt.Println(v1)
}

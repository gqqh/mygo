package main

import "fmt"

func main() {
	v1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(v1)
	var index int
	index = 1
	v1 = append(v1[:index-1], v1[index+1:]...)
	fmt.Println(v1, "= [1, 2, 3, 4, 5]")

	index = 4
	v1 = append(v1[:index-1], v1[index+1:]...)
	fmt.Println(v1, "= [1, 2, 3, 4]")

	index = 2
	v1 = append(v1[:index-1], v1[index+1:]...)
	fmt.Println(v1, "= [1, 2, 4, 5]")
}

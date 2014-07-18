// slice.go
package main

import (
	"fmt"
)

func main() {
	/*	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println(a)
		s1 := a[5:] //s1 := a[5:10] or s1:=a[5:len(a)]
		s2 := a[:5]
		fmt.Println(s1)
		fmt.Println(s2)
	*/
	//make slice
	/*	s1 := make([]int, 3, 10)
		fmt.Println(len(s1), cap(s1))
		fmt.Println(s1)
	*/

	/*	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
		slice_a := a[2:5]
		slice_b := a[3:5]
		fmt.Println(string(slice_a))
		fmt.Println(string(slice_b))
		a[4] = '!'
		fmt.Println(string(slice_a))
		fmt.Println(string(slice_b))

		sc := slice_a[2:6]
		fmt.Println(string(sc))
		sc = slice_a[2:]
		fmt.Println(string(sc))
	*/
	/*
		s1 := make([]int, 3, 6)
		fmt.Printf("%p\n", s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p\n", s1, s1)
		s1 = append(s1, 1, 2, 3)
		fmt.Printf("%v %p\n", s1, s1)
	*/

	/*
		a := []int{1, 2, 3, 4, 5}
		s1 := a[2:5]
		s2 := a[1:3]
		fmt.Println(s1, s2)
		s1[0] = 9
		fmt.Println(s1, s2)                    //s1, s2指向同一个底层数组
		fmt.Println(len(s2), cap(s1), cap(s2)) //cap()求得是从其索引开始到底层数组末端的长度
		//当s2增加元素超过其cap()时，会另起炉灶新建一个底层数组。
		s2 = append(s2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1)
		s1[0] = 10
		fmt.Println(s1, s2)
	*/
	/*
		s1 := []int{1, 2, 3, 4, 5, 6}
		s2 := []int{7, 8, 9}
		fmt.Println(s1, s2)
		fmt.Println(len(s1), cap(s1))
		fmt.Println(len(s2), cap(s2))
		copy(s1, s2)
		fmt.Println(s1, s2)
		fmt.Println(len(s1), cap(s1))
		fmt.Println(len(s2), cap(s2))

		copy(s2, s1)
		fmt.Println(s1, s2)
		fmt.Println(len(s1), cap(s1))
		fmt.Println(len(s2), cap(s2))
	*/
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	fmt.Printf("%p %p", s1, s2)
	fmt.Println(s1, s2)
	s2 = append(s2, 0)
	fmt.Printf("%p %p", s1, s2)
	fmt.Println(s1, s2)

	s3 := s1[:]
	fmt.Printf("%p %p %p", s1, s2, s3)
	fmt.Println(s1, s2, s3)
}

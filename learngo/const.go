// const.go
package main

import (
	"fmt"
)
/*
const a int = 1
//b := 'A'

const (
	b = a
	c = a + 1
)*/
const (
	a = 'A'
	b = iota
	c = 'B'
	d = iota
)
const (
	e = iota
)
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	
}

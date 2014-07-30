package main

import (
	"fmt"	
)
/*
const (
	_ = iota
	B
	KB float64 = 1 << (iota * 10)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB 
)

func main(){
	fmt.Println(B)
	fmt.Println("KB = ", KB)
	fmt.Println("MB = ", MB)
	fmt.Println("GB = ", GB)
	fmt.Println("TB = ", TB)
	fmt.Println("PB = ", PB)
	fmt.Println("EB = ", EB)
	fmt.Println("ZB = ", ZB)
	fmt.Println("YB = ", YB)
}*/

func main(){
	a := 1
	a ++
	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)
}
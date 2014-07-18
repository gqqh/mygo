package main

import (
	"fmt"
	"time"
)

func main(){
	for i:=0; i < 10; i++{
		go Add(i, 8)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("over...")
}

func Add(i, j int) int {
	fmt.Println(i+j)
	return i + j
}

package main

import "fmt"

func Count(cha chan int){
	fmt.Println("Counting...")
//cha <- 1
	<- cha
}

func main(){
	chas := make([]chan int, 10)
	for i:= 0; i<10; i++{
		chas[i] = make(chan int)
		go Count(chas[i])
	}

	for _, ch:= range(chas){
		ch <- 1
	}
	fmt.Println("over...")
}

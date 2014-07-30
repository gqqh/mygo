package main

import "fmt"

func adder() func(int) int { //闭包累加
	sum := 0 //闭包的全局变量,每个闭包公用一个
	fmt.Println(&sum)
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() //实例化两个闭包
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// defer_panic.go
package main

import (
	"fmt"
)

func main() {
	/*	A()
		B()
		C()*/
	var fs = [4]func(){}
	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i =", i)                       //直接值拷贝，传参
		defer func() { fmt.Println("defer_closure i = ", i) }() //匿名函数，均为闭包，i为地址引用
		fs[i] = func() { fmt.Println("closure i = ", i) }       //匿名函数，均为闭包，i也是地址引用
	}
	for _, f := range fs {
		f()
	}
	//输出如下
	/*
		closure i =  4
		closure i =  4
		closure i =  4
		closure i =  4
		defer_closure i =  4
		defer i = 3
		defer_closure i =  4
		defer i = 2
		defer_closure i =  4
		defer i = 1
		defer_closure i =  4
		defer i = 0
	*/
}

/*
func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() { //defer 一个匿名函数
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("panic in B")
	fmt.Println("Func B") //永远不会执行
}

func C() {
	fmt.Println("Func C")
}
*/

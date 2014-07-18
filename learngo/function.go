// function.go
package main

import (
	"fmt"
	//	"math"
)

/*
	Go函数不支持嵌套、重载和默认参数
	支持：
		无需声明原型、不定长度变参、多返回值、命名返回值参数、
		匿名函数、闭包
	定义函数使用关键字func，且左大括号不能另起一行
	函数也可以作为一种类型使用
*/
func main() {
	/*	a, b, c := 1, 2, 3
		//	C(a, b, c)	//值传递
		//C(&a, &b, &c)	//指针传递
		s1 := []int{a, b, c} //slice是地址操作，故为指针传递
		C(s1)
		fmt.Println(s1)
	*/
	//匿名函数
	/*
		a := func() {
			fmt.Println("func A")
		}
		a()
	*/
	//闭包
	/*	f := closure(10, 1)
		fmt.Println(f(0))
		fmt.Println(f(1))
		fmt.Println(f(2))
		fmt.Println(f(3))
	*/
	//defer
	/*for i := 0; i < 3; i++ {
		defer fmt.Println(i)	//defer语句在函数返回时调用
	}
	*/
	
	/*for i := 0; i < 3; i++ {
		/*	defer func() { //同a := func(){fmt.Println(i)}
			fmt.Println(i)
		}()*/
	/*	a := func() {
			fmt.Println(i)
		}
		defer a()	//3 3 3在函数退出时调用*/
	//}
}

/*
func A(a, b, c int, d string) (int, string, int) {
	return
}*/
/*
func B() (a, b, c int) { //此处已经声明了a,b,c，如果下面用a, b ,c :=1, 2 ,3则错误
	a, b, c = 1, 2, 3
	return //一般会使用 return a, b, c
}*/
//指针传递
/*func C(s ...*int) {
	*s[0] += 10
	*s[1] += 10
	fmt.Println(*s[0], *s[1], *s[2])
}*/
//slice地址传递
/*
func C(s []int) {
	s[0] += 11
	s[1] += 11
	fmt.Println(s)
}
*/
//闭包函数
/*
func closure(x, y int) func(int) int {
	fmt.Printf("%p %p\n", &x, &y)
	return func(a int) int {
		fmt.Printf("%p %p\n", &x, &y)
		return int(math.Pow(float64(x), float64(a)) + float64(y))
	}
}
*/
/*
defer
1）执行方式类似其他语言中的析构函数，在函数执行结束后按照调用顺序
的相反顺序逐个执行
2）即使函数发生严重错误时也会执行
3）支持匿名函数的调用
4）常用于资源清理、文件关闭、解锁以及记录时间等操作
5）通过与匿名函数配合可在return之后修改函数计算结果
6）如果函数体内某个变量作为defer时匿名函数的参数，则在定义defer
时即已经获得了拷贝，否则则是引用某个变量的地址
7）Go没有异常机制，但有panic/recover模式来处理错误
8）panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/

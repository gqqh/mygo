// struct.go
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
	/*Contact struct {
		Phone, City string
	}*/
}

//结构组合
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
	Sex  int
}
type student struct {
	human
	Name string
	Age  int
}

/*
type person1 struct {
	Name string
	Age  int
}
*/
func main() {
	//a,b,c三种初始化方式
	/*a := person{}
	a.Name = "joy"
	a.Age = 21
	b := person{"joee", 20}
	c := person{
		Name: "joeee",
		Age:  22,
	}
	fmt.Println(a, b, c)

	fmt.Println(a)
	A(a) //值传递
	fmt.Println(a)
	B(&a)          //地址传递
	fmt.Println(a) //{joy 100}

	//更常用的是下面的
	d := &person{
		Name: "jooe",
		Age:  21,
	}
	fmt.Println(*d)
	B(d)
	fmt.Println(*d)

	//匿名结构
	a := &struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(*a)

	//匿名结构的嵌套
	a := person{
		Name: "joy",
		Age:  21,
	}
	//匿名结构只能这么初始化
	a.Contact.Phone = "123231231"
	a.Contact.City = "hefei"
	fmt.Println(a)
	//匿名字段，可以不指定字段的名字，在初始化时必须严格按照顺序依次初始化
	//结构可以在同类型之间赋值，和比较
	a := person{Name: "joe", Age: 21}
	b := person1{Name: "joe", Age: 21}
	c := person{Name: "joe", Age: 21}
	fmt.Println(a == b) //invalid operation: a == b (mismatched types person and person1)
	fmt.Println(a == c)
	*/
	//组合方式的初始化
	a := teacher{Name: "joe", Age: 19, human: human{0}, Sex: 4}
	b := student{Name: "joe", Age: 21, human: human{1}}
	a.Name = "joe2"
	a.Age = 13
	//下面两种都可以
	a.Sex = 2 //当嵌入的字段与外层有同名字段时，默认的字段就外层的，内层的必须指明具体结构
	a.human.Sex = 11
	b.human.Sex = 3
	fmt.Println(a, b)
}

/*
func A(per person) {
	per.Age = 100
	fmt.Println(per)
}
func B(per *person) {
	per.Age = 100     //注意此处并不需要:*per.Age
	fmt.Println(*per) //{joy 100}
}
*/

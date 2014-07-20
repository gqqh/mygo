package main

//method.go
import (
	"fmt"
)

/*方法method
1）go中没有class，但依旧有method
2）通过显示说明receiver来实现与某个类型的组合
3）只能为同一个包中的类型定义方法，不能跨包定义
4）Receiver可以是类型的值或者指针
5）不存在方法重载
6）可以使用值或者指针来调用方法，编译器会自动完成转化
7）从某种意义上来说，方法是函数的语法糖，因为receiver其实就是方法所接收的
  第一个参数（Method value vs mathod expression）
8）如果外部结构和嵌套结构存在同名方法，则优先调用外部方法
9）类型别名不会拥有底层类型所附带的方法
10）方法可以调用结构中的非公开字段
*/

type A struct {
	Name string
}

type B struct {
	Name string
}

//类型别名
type TZ int //TZ是int的别名

func main() {
	/*	a := A{"a"}
		a.Print()
		fmt.Println(a.Name)

		b := B{"b"}
		b.Print()
		fmt.Println(b.Name) //修改失败
	*/
	var a TZ //从TZ类型转换成int类型时，也需要强制转换
	a = 0
	a.Print() //任何一种类型都可以添加方法 method value
	a.Increase()
	(*TZ).Print(&a) //method expression
}
func (a *A) Print() { //第一个参数一定是接收者，指针传递
	a.Name = "AA"
	fmt.Println("func A")
}

func (b B) Print() { //第一个参数一定是接收者, 值传递
	b.Name = "BB"
	fmt.Println("func B")
}

func (a *TZ) Print() {
	fmt.Println(*a)
}
func (a *TZ) Increase() {
	*a += 100
}

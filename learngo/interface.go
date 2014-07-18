// interface.go
package main

import (
	"fmt"
)

/*接口interface
1）接口是一个或多个方法签名de集合
2）只要某个类型拥有该接口的所有方法签名，即算实现该接口，无需显示声明实现了哪个接口
  这称为Structal Typing
3）接口只有方法声明，没有实现，没有数据字段
4）接口可以匿名嵌入其他接口，或嵌入到结构中
5）将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，
  既无法修改复制品的状态，也无法获取指针
6）只有当接口存储的类型和对象都为nil时，接口才等于nil
7）接口调用不会做receiver的自动转换
8）接口同样支持匿名字段方法
9）接口也可实现类似OOP中的多态
10）空接口可以作为任何类型数据的容器
11）接口的转换只能由父类转换成子类
*/
//空接口，相当于Java里面的Object
type emptyer interface{}

//定义一个接口
type USBer interface { //一般命名接口都以er结尾
	Name() string //method
	//	Connect()
	Connecter //嵌入接口
}

//Connecter接口
type Connecter interface {
	Connect()
}

//定义一个结构
type PhoneConnecter struct {
	name string
}

func main() {
	var a USBer
	a = PhoneConnecter{"a phoneconnecter"}
	a.Connect()
	DisConnect(a)
}

//PhoneConnecter实现接口USBer的两个方法,就算是实现了该接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

/*
//判断一个变量是什么类型，不能判断是否为一个接口
func DisConnect(usb USBer) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected:", pc.name)
		return
	}
	fmt.Println("Unknown device")
}
*/
//使用空接口的方法
func DisConnect(usb interface{}) { //传入一个空接口
	switch v := usb.(type) { //让系统猜usb的类型,注意此处只有一个返回值
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown device...")

	}
}

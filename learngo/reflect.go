// reflect.go
package main

import (
	"fmt"
	"reflect" //reflect包
)

/*反射reflection
1）反射可以大大提高程序的灵活性，使得interface{}有更大的发挥余地
2）反射使用TypeOf和ValueOf函数从接口中获取目标对象信息
3）反射会将匿名字段作为独立字段（匿名字段本质）
4）想要利用反射修改对象状态，前提是interface.data是settable，
  即pointer-interface
5）通过反射可以“动态”调用方法
*/

//定义一个结构
type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User  //嵌套结构
	title string
}

func main() {
	/*u := User{1, "yoe", 23}
	Info(u) //此处为值传递
	//Info(&u) //此处为传递一个指针，不是一个reflect.struct
	m := Manager{User: User{1, "joy", 21}, title: "1234"}
	t := reflect.TypeOf(m) //获取m的类型信息
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	x := 123
	v := reflect.ValueOf(&x)
	fmt.Println(x)
	v.Elem().SetInt(999) //通过反射修改结构的值,必须是传递指针
	fmt.Println(x)

	u := User{1, "joe", 12}
	Set(&u)
	fmt.Println(u)
	*/

	//通过反射动态调用方法
	u := User{1, "OK", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joee"), reflect.ValueOf(23)} //设置参数，必须为一个slice

	re := mv.Call(args) //调用
	fmt.Println(re)
}

/*
func Set(obj interface{}) {
	v := reflect.ValueOf(obj)

	//判断v是否是一个reflect.Ptr，以及v的元素是否是可以修改的
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem()
	}

	//如果判断是否找到名字
	if f := v.FieldByName("Name"); !f.IsValid() {
		fmt.Println("cannot find field 'Name'.")
		return
	}
	//使用名字查找域
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("hhhha")
	}
}
*/
//通过反射动态调用
func (user User) Hello(name string, age int) (an string) {
	fmt.Println("Hello", name,
		", my name is", user.Name,
		"\nmy age is", age)
	an = "all right"
	return
}

/*
//使用reflect方法
func Info(obj interface{}) {
	t := reflect.TypeOf(obj) //获取obj的类型信息
	fmt.Println("Type:", t.Name())

	//判断传的参数是否是我们需要的类型
	if k := t.Kind(); k != reflect.Struct { //如果k不是一个反射结构，则返回
		fmt.Println("XXX")
		return
	}

	v := reflect.ValueOf(obj) //获取obj的值信息
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}
	//打印方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s:%v\n", m.Name, m.Type)
	}
}
*/

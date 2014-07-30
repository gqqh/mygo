//slice坑
package main

import (
	"fmt"
)

func add(t []int) []int {
	t = append(t, 10)
	//当slice增加，且大小大于底层的数组大小时，会重新开辟一个空间，并把现在的都复制过去，但是，返回地址并不修改
	return t //由于新产生了一个底层数组，所以要通过显示返回
}

func main() {
	t := make([]int, 0) //size 0, slice
	fmt.Println(t)
	t = add(t)
	fmt.Println(t)
}

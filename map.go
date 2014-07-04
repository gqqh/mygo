// map.go
package main

import (
	"fmt"
	//	"sort"
)

func main() {
	//初始化map
	/*var m map[int]string
	m = make(map[int]string, 10)	//10是初始大小，可忽略
	m = {1:"ok"}*/
	/*m := make(map[int]string, 10)
	m = {1:"ok"}*/
	/*m := map[int]string{1: "ok"}
	fmt.Println(m[1])
	*/
	//map基本用法
	/*
		var m map[int]map[int]string
		m = make(map[int]map[int]string)
		m[1] = make(map[int]string)
		m[1][1] = "ok"
		a, ok := m[2][1]	//两个返回值， ok为是否读取成功标志
		if !ok {
			m[2] = make(map[int]string)
			m[2][1] = "Good"
		}
		a, ok = m[2][1]
		fmt.Println(a, ok)
		a, ok = m[1][1]
		fmt.Println(a, ok)
	*/
	//for结合slice的遍历使用方法
	/*	s1 := []int{1, 2, 3, 4}
		for i, v := range s1 { //i=index, v = slice[i], v是拷贝修改不影响原来的值
			v += 10
			fmt.Println(i, v)
		}
		fmt.Println(s1)
		//如果需要修改原来的值，则需要使用索引直接操作
		for i := range s1 { //i=index, 直接修改s1[i]
			s1[i] += 10
			fmt.Println(i, s1[i])
		}
		fmt.Println(s1)
	*/
	//for结合map的遍历使用方法
	/*	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
		for k, v := range m { //k=key, v=map[k], 也是拷贝
			v += "!"
			fmt.Println(k, v)
		}
		fmt.Println(m)
		for k := range m { //k=key, 直接修改map[k]
			m[k] += "!"
			fmt.Println(k, m[k])
		}
		fmt.Println(m)
	*/
	//slice和map结合使用，slice的每个元素均为map
	/*
		sm := make([]map[int]string, 5, 10) //slice
		for i, v := range sm {
			v = make(map[int]string)
			v[i] = "ok"
			fmt.Println(i, v[i])
		}
		fmt.Println(sm)

		for i := range sm {
			sm[i] = make(map[int]string)
			sm[i][i] = "ok"
			fmt.Println(i, sm[i][i])
		}
		fmt.Println(sm)
	*/
	//map的key值是无序的，需要结合slice排序
	/*m := map[int]string{1: "a", 4: "d", 2: "b", 5: "e", 6: "g", 3: "c"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)*/
	//map的键和值互换
	m1 := map[int]string{1: "a", 4: "d", 2: "b", 5: "e", 6: "g", 3: "c"}
	m2 := make(map[string]int, len(m1))
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m1)
	fmt.Println(m2)
}

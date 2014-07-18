// gorountine.go
package main

import (
	"fmt"
	//	"time"
	"runtime"
	"sync"
)

/*并发concurrency
1）gorountine只是由官方实现的超级”线程池“而已，但是每个实例只有4-5KB的栈内存占用和
  由于实现机制而大幅减少的创建和销毁的开销，是制造Go高并发的原因。
2）并发不是并行，Go可以设置使用的核数，以发挥多核计算机的能力
3）goroutine奉行通过通信来共享内存，而不是共享内存来通信
*/
/*Channel
1）Channel是goroutine沟通的桥梁，大都是阻塞同步的
2）通过make创建，close关闭
3）Channel是引用类型
4）可以使用for range来迭代不断操作channel
5）可以设置单向或者双向通道
6）可以设置缓存大小，在未被填满前不会发生阻塞
*/
func main() {
	/*chanl := make(chan bool) //创建一个布尔型通道，为双向的，无缓冲的，会同步阻塞

	go func() {
		fmt.Println("匿名goroutine!!!")
		chanl <- true
	}()
	<-chanl
	
	go func() {
		fmt.Println("Go go go...")
		chanl <- true
		close(chanl)
	}()
	//如果以上不关闭的话，则提示：all goroutines are asleep - deadlock!
	for v := range chanl {
		fmt.Println(v)
	}
	//	go Go()
	//time.Sleep(2 * time.Second)
	*/
	/*
	//设置CPU，通过设置带缓存的channel来同步
	runtime.GOMAXPROCS(runtime.NumCPU())	//设置CPU个数
	c := make(chan bool, 10)	//设置带缓存的channel
	for i:= 0; i < 10; i++{
		go Goo(c, i)
	}
	for i:=0; i< 10; i++{
		<- c
	}
	*/
	//通过sync包的工作组解决同步问题
	runtime.GOMAXPROCS(runtime.NumCPU()) 	//设置CPU个数
	runtime.GOMAXPROCS(5) 	//设置CPU个数，当设置的CPU核数大于系统的核数时，只会按最大的运行
	wg := sync.WaitGroup{} 	//设置同步的等待同步组
	wg.Add(10) 	//设置该组有10个成员
	for i:= 0; i<10; i++{
		go Go(&wg, i)	//由于Go是值传递，而wg变量需要值修改，所以传的是指针
	}
	
	wg.Wait()	//等待所有goroutine执行完	
}
//通过sync包的工作组解决同步问题
func Go(wg *sync.WaitGroup, index int){
	a := 1
	for i:=0; i< 100000000000; i++{
		a += i
	}
	fmt.Println(index, a)
	wg.Done()	//通知工作组，该协程已完成，wg默认会减1
}
/*
//通过带缓存的channel来解决同步的问题
func Goo(c chan bool, index int){
	a := 1
	for i:=0; i < 10000000; i++{
		a += i
	}
	fmt.Println(index, a)
	c <- true
}
*/
/*
func Go() {
	fmt.Println("Go go go!!!")
}
*/

// gorountine.go
package main

import (
	"fmt"
	//	"time"
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
	chanl := make(chan bool) //创建一个布尔型通道，为双向的，无缓冲的，会同步阻塞

	/*go func() {
		fmt.Println("匿名goroutine!!!")
		chanl <- true
	}()
	<-chanl
	*/
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
}

func Go() {
	fmt.Println("Go go go!!!")
}

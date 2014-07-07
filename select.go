//select.go
package main

import "fmt"
/*Select
1）可处理一个或多个channel的发送与接收
2）同时有多个可用的channel时按随机顺序处理
3）可用空的select来阻塞main函数
4）可设置超时
*/

func main(){
  //使用select接收channel
/*  c1, c2 := make(chan int), make(chan string)	//创建两个channel
  o := make(chan bool, 2)	//输出标志
  go func(){
    for{	//while(true)
      select{	
        //select的case与switch的不同之处在于select的case条件必须是一个channel的读写
        case v, ok := <-c1:
          if !ok{
            fmt.Println("c1")
            o <- true
            break
          }
          fmt.Println("c1", v)
        case v, ok := <-c2:
          if !ok{
            fmt.Println("c2")
            o<-true
            break
           }
           fmt.Println("c2", v)
      }
    }
  }()
    
  c1 <- 1
  c2 <- "Hi"
  c1 <- 3
  c2 <- "hello"
    
  close(c1)
  close(c2)
    
  for i:=0; i<2;i++{
    <-o
  }
  */
  //使用select发送channel
 /* c:=make(chan int)
  go func(){
    for v:=range c{
      fmt.Println(v)
    }
  }()
  
  for i:=0; i<100; i++{
    select{
      case c <- 0:
      case c <- 1:
    }
  }*/
  //select{} 	//可用阻塞主程序
  
  
}

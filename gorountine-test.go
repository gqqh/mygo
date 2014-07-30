//goroutine test
package main

import (
	"fmt"
)

var ch chan string

func Play(str string) {
	i := 0
	for {
		fmt.Println(<-ch)
		ch <- fmt.Sprintf("From %s: Hi, #%d", str, i)
		i++
	}
}

func main() {
	ch = make(chan string)
	go Play("ping")
	//go Play("pong")
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("From main: Hello, #%d", i)
		fmt.Println(<-ch)
	}
}

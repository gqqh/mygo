//gouroutine_test.go
package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball) //game on; toss the ball
	time.Sleep(1 * time.Second)
	<-table //game over; grab the ball

	panic("show me the stack")
}

func player(str string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(str, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

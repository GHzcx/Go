package main

import (
	"fmt"
	"time"
)

func main () {
	Queue := make(chan int, 10)
	for i :=0; i < 100 ; i++ {
		Queue <- 1
		go func(num int, Queue chan int) {
			fmt.Println(num)
			time.Sleep(1 * time.Second)
			<-Queue
		}(i, Queue)
		fmt.Println("Queue", len(Queue))
	}
}
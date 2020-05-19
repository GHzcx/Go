package main

import (
	"fmt"
	"time"
)
func test(sig chan int) {
	sig <- 1
}

func main () {
	//Queue := make(chan int, 10)
	//for i :=0; i < 100 ; i++ {
	//	Queue <- 1
	//	go func(num int, Queue chan int) {
	//		fmt.Println(num)
	//		time.Sleep(1 * time.Second)
	//		<-Queue
	//	}(i, Queue)
	//	fmt.Println("Queue", len(Queue))
	//}
	sig := make(chan int)
	go test(sig)
	select {
	case s := <-sig:
		fmt.Println("Signal", s)
	default:
		time.Sleep(3 * time.Second)
		fmt.Println("default")
	}
}
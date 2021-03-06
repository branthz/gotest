package main

import (
	"fmt"
	"time"
)

func main() {
	go doselect()
	c := make(chan int)
	<-c
}

func doselect() {
	ch := make(chan int)
	ch2 := make(chan int, 10)
	t := time.NewTicker(time.Second * 10)
	select {
	case <-ch:
		fmt.Println("send ch")
	case ch2 <- 10:
		fmt.Println("recv ch")
	case <-t.C:
		fmt.Println("ticker!")
		//default:
		//	fmt.Println("execute default")
	}
}

func recv() int {
	return 10
}

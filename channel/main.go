package main

import (
	"fmt"
	//"os"
	"sync"
	"time"
)

var datach chan int
var sig chan bool

type kvstore struct {
	val     int
	mu      sync.RWMutex
	kvStore map[string]string // current committed key-value pairs
}

func main(){
	ch :=make(chan int,10)
	ch <- 1
	ch <- 2
	close(ch)
	for v:= range ch{
		fmt.Println(v)
	}
}

func xxx() {
	sc, ec := newchannel()
	/*
		go func() {
			if err, ok := <-errorC; ok {
				fmt.Println(err)
			}
			fmt.Printf("game over\n")
			os.Exit(0)
		}()*/

	ob := newkvstore(sc, ec)
	fmt.Printf("=======%d\n", ob.val)

	time.Sleep(1e9 * 10)

	//stop := make(chan int)
	//<-stop
}

func newchannel() (<-chan *string, <-chan error) {
	strch := make(chan *string)
	errorC := make(chan error)
	//go writeChannel(strch)
	return strch, errorC
}

func writeChannel(ch chan<- *string) {
	var str = "nihao"
	for {
		ch <- &str
		time.Sleep(1e9 * 2)
	}
}

func newkvstore(c <-chan *string, e <-chan error) *kvstore {
	s := &kvstore{val: 20, kvStore: make(map[string]string)}

	s.readCommits(c, e)
	fmt.Printf("55555555\n")
	go s.readCommits(c, e)
	return s
}

func (s *kvstore) readCommits(commitC <-chan *string, errorC <-chan error) {
	fmt.Printf("999999999\n")
	for data := range commitC {
		fmt.Printf("-----lalal\n")
		if data == nil {
			fmt.Printf("1111111\n")
			return
		}
		fmt.Printf("%s\n", *data)
	}

	if err, ok := <-errorC; ok {
		fmt.Println(err)
	}
}

func benchChannelBuffer() {
	var num = 100

	datach = make(chan int, 10000)
	go getdataex(0)
	for i := 1; i < num; i++ {
		go getdata(i)
	}

	for j := 0; j < 5; j++ {
		for i := 200; i < 300; i++ {
			datach <- i
		}
		time.Sleep(1e8)
	}
	time.Sleep(1e9)
}
func getdataex(index int) {
	for {
		t := <-datach
		//time.Sleep(1e8 * time.Duration((101-index)%10))
		fmt.Printf("go routine:%d get data:%d\n", index, t)
		<-sig
	}
}
func getdata(index int) {
	for {
		t := <-datach
		//time.Sleep(1e8 * time.Duration((101-index)%10))
		fmt.Printf("go routine:%d get data:%d\n", index, t)
	}
}

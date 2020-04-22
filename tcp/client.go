package main

import (
	"fmt"
	"net"
)

func send() {
	raddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:9800")
	if err != nil {
		fmt.Println(err)
		return
	}
	tp, err := net.DialTCP("tcp4", nil, raddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf = make([]byte, 1024)
	for {
		_, err = tp.Read(buf)
		if err != nil {
			fmt.Printf("++++%v", err)
			break
		}
	}
	tp.Close()
}

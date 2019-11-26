package main

import (
	"fmt"
	"net"
)

func send() {
	raddr, err := net.ResolveTCPAddr("tcp4", "192.168.29.100:9800")
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
			break
		}
	}
	tp.Close()
}

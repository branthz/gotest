package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

var port int = 9800

func udpServ() error {
	addr, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return err
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var data = make([]byte, 1460)
	var rn int
	for {
		rn, err = conn.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("conn read:%s\n", string(data[:rn]))
	}
	fmt.Println("udp server quit\n")
	return nil
}

func tcpserv() {
	addr, err := net.ResolveTCPAddr("tcp4", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}
	ln, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		fmt.Println("listen:", err)
		return
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Println("accept:", err)
			return
		}
		go handleConnection(conn)
	}

}

var tp string

func main() {
	flag.IntVar(&port, "p", 9800, "listen port")
	flag.StringVar(&tp, "t", "udp", "choose udp or tcp")
	flag.Parse()
	if tp == "udp" {
		udpServ()
	} else {
		tcpserv()
	}
}

func handleConnection(c *net.TCPConn) {
	c.Write([]byte("ok"))
	c.SetReadDeadline(time.Now().Add(time.Second * 3))
	c.Close()
	return
}

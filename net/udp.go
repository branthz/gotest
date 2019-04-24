package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	outAddr *string = flag.String("out", "112.124.35.104:4998", "outside listen port")
	inAddr  *string = flag.String("in", "127.0.0.1:4999", "inner listen port")
)

func udpServer() {
	flag.PrintDefaults()
	flag.Parse()
	oaddr, err := net.ResolveUDPAddr("udp4", *outAddr)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	iaddr, err := net.ResolveUDPAddr("udp4", *inAddr)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	UConn, err := net.ListenUDP("udp4",
		&net.UDPAddr{
			Port: oaddr.Port,
		})
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}
	defer UConn.Close()

	for {
		buf := make([]byte, 1024)
		rn, from, err := UConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("read from %s:%v\n", from.String(), buf[:rn])
		UConn.WriteToUDP(buf[:rn], iaddr)
	}
}

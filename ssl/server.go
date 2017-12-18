package main

import (
	"crypto/tls"
	"fmt"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./server.pem", "./skey.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	cert2, err := tls.LoadX509KeyPair("./server2.pem", "./sker2.pem")
        if err != nil {
                fmt.Println(err)
                return
        }
	tc := &tls.Config{Certificates: []tls.Certificate{cert,cert2}}
	tc.BuildNameToCertificate()
	sl, err := tls.Listen("tcp", ":6666", tc)
	for {
		conn, err := sl.Accept()
		if err != nil {
			continue
		}
		go tcpHandle(conn)
	}

}

func tcpHandle(conn net.Conn) {
	var buf [128]byte
	rn, err := conn.Read(buf[:])
	if err != nil {
		fmt.Printf("====1%v\n", err)
		return
	}
	fmt.Println(string(buf[:rn]))
	_, err = conn.Write([]byte("world"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

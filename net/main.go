package main
import(
	"net"
	"fmt"
)

func main(){
	ln, err := net.Listen("tcp", "192.168.3.111:7979")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue	
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	var buf [1024]byte
	n,err:=conn.Read(buf[:])
	fmt.Printf("--------%d,%v----\n",n,err)
	conn.
}

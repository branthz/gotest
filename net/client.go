package main
import(
	"net"
	"fmt"
	"time"
)
func main(){
	conn, err := net.Dial("tcp", "192.168.3.111:7979")
	if err != nil {
		fmt.Printf("-----%v\n",err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	time.Sleep(1e9)
}

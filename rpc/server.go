package main
import(
	"net/rpc"
	"net"
	"./proto"
	"net/http"
)

func main(){
	c:=new(proto.Client)
	rpc.Register(c)
	rpc.HandleHTTP()

	l,_:=net.Listen("tcp",":1234")	
	http.Serve(l,nil)
}

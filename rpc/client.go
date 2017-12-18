package main
import(
	"net/rpc"
	"fmt"
	"./proto"
)

func main(){
	c,err:=rpc.DialHTTP("tcp","localhost:1234")
	if err!=nil{
		fmt.Println(err)
		return	
	}
	args:= &proto.Args{15,8}
	var res proto.Param
	err=c.Call("Client.Divide",args,&res)
	if err!=nil{
		fmt.Println(err)
		return	
	}
	fmt.Println(res.Yu,res.Chu)
}

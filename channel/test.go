package main
import(
	"fmt"
	"time"
)
func main(){
	c:=make(chan string)
	go func(){
		time.Sleep(1e9*3)
		close(c)
	}()
	//如果channel被关闭，则ok为false
	//if v,ok := <-c;ok{
	//	fmt.Printf("not blocked,%s\n",v)
	//}
	v := <-c
	fmt.Printf("end---%s\n",v)
}

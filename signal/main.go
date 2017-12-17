package main
import(
	"fmt"
	"os"
	"os/signal"
//	"runtime"
	"time"
)

func main(){
	sig := make(chan os.Signal,1)
	//done:= make(chan boll,1)
	signal.Notify(sig,os.Interrupt,os.Kill)

	s:= <-sig
	fmt.Printf("signal get:%v\n",s)
	
	time.Sleep(1e9*10)
}

package main
import(
	"time"
	"fmt"
)

func gettimediff() uint32 {
	timestamp1 := time.Now().Unix()
	//time2, _ := time.Parse("2006-01-02 15:04:05", "2001-03-04 05:06:07")
	diff := timestamp1 - 999999
	return uint32(diff)
}

func (recv uint32) Checktimediff() bool {
	a := recv - gettimediff()
	a = (a ^ a>>31) - a>>31
	if a > 30 {
		return false
	}
	return true
}
func main(){
	tm:=time.Now().UTC()
	x1,x2:=tm.Zone()
	fmt.Printf("%s--%d\n",x1,x2)
	fmt.Printf("%v==========\n",uint32(time.Now().Unix()).Checktimediff())
}

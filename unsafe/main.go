package main
import(
    "unsafe"
    "fmt"
)
type yy struct{
    a [32]byte
}
type xx struct{
    c [8]byte
    s [16]byte
    f [80]byte
    g yy
}

func main(){
    var a xx
    p:=unsafe.Sizeof(a)
    fmt.Printf("---%d-------\n",p)
    //size()
}

func size(){
	type hah struct {
		a uint32
		b uint16
		c [6]byte
	}

	//var ni hah
	fmt.Printf("----%d--\n",unsafe.Sizeof(hah{}))
}

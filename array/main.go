package main
import(
	"fmt"
)

func main(){
	var a = "杭州欢迎您"
	var b = []byte(a)
	b[0]='广'
	var c = string(b)
	fmt.Printf("%p----%p---%p\n",&a,b,&c)
	fmt.Printf("%v----%v---%v\n",a,b,c)
}

func copyNotPoint(){
	var a=[3]int{1,2,3}
	add(a)
	fmt.Printf("%v\n",a)	
}

func add(a [3]int){
	for i:=0;i<len(a);i++{
		a[i]+=1
	}
}


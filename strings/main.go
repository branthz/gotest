package main
import(
	"fmt"
	"strings"
)
func trim(){
	var str = "*   *hello *&^%$ world *     "
	s:=strings.Trim(str,"*    ")
	fmt.Printf("%s---%d\n",s,len(s))
}

func main(){
	var a = []int{1,2,3,4,5}
	change(a)
	fmt.Println(a)
	var b =make(map[string]interface{})
	b["hello"]="shan"
	mchange(b)
	fmt.Println(b)
}

func change(s []int){
	//s=[]int{9,9,8}
	s[2]=10
	return
}

func mchange(m map[string]interface{}){
	//m["hello"]="world"
	m=make(map[string]interface{})
	m["ni"]="hao"
	return
}

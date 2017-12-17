package main
import(
	"fmt"
	"container/list"
)

func main(){
	l := list.New()
	e4 := l.PushBack(4)
	//e1 := l.PushFront(1)
	//l.InsertBefore(3, e4)
	//l.InsertAfter(2, e1)
	e5 :=l.PushBack(5)
// Iterate through list and print its contents.
//for e := l.Front(); e != nil; e = e.Next() {
//    fmt.Println(e.Value)
//}
	l.Remove(e4)
	l.Remove(e5)
	l.PushBack(6)
	var pl *int
	pl=nil
	l.PushBack(pl)
	fmt.Println(l.Len())
}

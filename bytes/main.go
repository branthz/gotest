package main

import (
	"bytes"
	"fmt"
)

type ByteView struct {
	// If b is non-nil, b is used, else s is used.
	b []byte
	s string
}

func (v ByteView) Slice(from, to int) ByteView {
	if v.b != nil {
		return ByteView{b: v.b[from:to]}
	}
	return ByteView{s: v.s[from:to]}
}
func haha() {
	var xx = ByteView{
		s: "123456",
	}
	h := xx.Slice(2, 4)
	fmt.Printf("%v\n", h)
	str := "987654321"
	fmt.Printf("%v\n", str[2:6])
}

func main(){
	//test if []byte(string) share the same memory with string
	var a = []byte("hello")
	var b =string(a)
	a[2]='c'
	var c = []byte(b)
	c[2]='q'
	fmt.Println(string(a),b)
}


func testBytesBuffer() {
	tbuf := bytes.NewBuffer(make([]byte, 0))
	var val int = 25
	fmt.Fprintf(tbuf, "a*%10d\r\n", val)
	tslice := tbuf.Bytes()
	var x string = fmt.Sprintf("%10d", 123456)
	copy(tslice[2:11], []byte(x))

	fmt.Printf("%d---%s----%v\n", tbuf.Len(), tbuf.String(), tslice)

	fmt.Printf("%v\n",string(65))
}

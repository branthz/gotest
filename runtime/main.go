package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"unicode"
)

/*
func runChild() {
	var args []string
	args = append(args, os.Args[1:]...)
	args = append(args, "-child")
	cmd := exec.Command(os.Args[0], args...)
	fmt.Printf("haha\n")

}
*/
func cmd() {

	fmt.Printf("%d\n", os.Getppid())
	/*err := runChild()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	*/
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		fmt.Printf("=======%v\n", err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	time.Sleep(1e9 * 10)
	return
}
func main(){
	//testCopy()
	testRune()
}

func testRune(){
	var a="123abc"
	if unicode.IsDigit(rune([]byte(a)[0])){
		fmt.Println("digit")
	}else{
		fmt.Println("not digit")
	}
}

const copyCount = 100*1000
func testCopy(){
	var a [1024]byte
	var b [1024]byte
	for i:=0;i<1000;i++{
		b[i]=byte(i)
	}
	ts:=time.Now().UnixNano()
	for i:=0;i<copyCount;i++{
		copy(a[:],b[:1000])
	}
	tp:=time.Now().UnixNano()
	diff:=(tp-ts)/1e6
	fmt.Printf("time diff:%d\n",diff)
}


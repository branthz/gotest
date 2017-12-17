package main

import (
	"bytes"
	"fmt"
	"log"
	xx "package/log"
	"unsafe"
	"runtime"
	"time"
)

type dd struct{
	name [30]byte
}
var(
	mlog       *xx.Logger
)
func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	//fmt.Print(&buf)
	//mlog, _ = xx.New("", "DEBUG")
	//mlog.Error("nihao")
	var a [10]byte
	a[5]=10
	var b int=10

	var ph =(*dd)(unsafe.Pointer(&a[0]))
	var format  string="hahahah---,%v,%s\n"
	out:=fmt.Sprintln(format,a,b,ph)
	//out:=fmt.Sprintln(format,"hello","world")
	runtime.GOMAXPROCS(runtime.NumCPU())
	//out :="hello world"
	for i:=0;i<100;i++{
		go multiLogging(out)
	}
	time.Sleep(1e9*60)
	//fmt.Printf("out:%s\n",out)
	//mlog.Debugln("---%s",out,b)
}

func multiLogging(h string){
	for i:=0;i<100;i++{
		mlog.Debugln(h)
	}
}


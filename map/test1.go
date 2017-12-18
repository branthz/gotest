package main

import (
	"fmt"
	"time"
	"unsafe"
)

type addr struct {
	a1 string
	b1 string
}

func getone() {
	xx := make(map[string]*addr)
	var t1 addr
	var t2 addr
	t1.a1 = "111"
	t1.b1 = "222"
	xx["11"] = &t1
	t2.a1 = "333"
	t2.b1 = "444"
	xx["22"] = &t2
	delete(xx, "11")
	fmt.Println(xx, t1)
}

const REPEAT = 10000000

func main() {
	//getone()
	//keyPerformance()
	doubleMap()
}

type mmap map[string]interface{}

//test doubel map function
func doubleMap() {
	//xx := make(map[string]interface{})
	//xx["hello"] = "world"
	//(xx["pepole"].(mmap))["zhangsan"] = 20
	var a = make(mmap)
	var b = make(mmap)
	b["hello"] = "world"
	a = b
	a.(mmap)["zhangsan"] = 21
	c := a.(mmap)
	fmt.Println(c)
}

func keyPerformance() {
	//ret:=getone()
	//fmt.Printf("%s\n",ret)
	//map_copy()
	var mac = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	var macs = mac[:]
	//var session = []byte{0x11, 0x22, 0x22, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99,0xaa,0xbb,0xcc,0xee,0xdd,0xff,0x00}
	var session = [16]byte{0x11, 0x22, 0x22, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xee, 0xdd, 0xff, 0x00}
	var p = (*int32)(unsafe.Pointer(&mac[0]))
	pp := make(map[string][16]byte, REPEAT)
	//pp := make(map[[32]byte][16]byte, REPEAT)
	ts := time.Now().UnixNano() / 1e6
	for i := 0; i < REPEAT; i++ {
		pp[string(macs)] = session
		//pp[mac] = session
		(*p)++
	}
	var tlen int = 0
	for k, _ := range pp {
		tlen += len(k)
	}
	tp := time.Now().UnixNano() / 1e6
	fmt.Printf("lalalalala----------------%d----%d\n", tp-ts, tlen)
}

func write100wan() {

}
func map_copy() {
	xx := make(map[int]addr)
	var t1 addr
	t1.a1 = "111"
	t1.b1 = "222"
	xx[1] = t1

	t1.a1 = "333"
	xx[2] = t1

	for k, v := range xx {
		fmt.Printf("---k:%d----v:%v\n", k, v)
	}
}

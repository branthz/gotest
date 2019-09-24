package main

import (
	"fmt"
	//"gopkg.in/redis.v2"
	"encoding/binary"
	"os"
	"time"
	"unsafe"

	"github.com/garyburd/redigo/redis"
	//"reflect"
)

type session struct {
	a int32
	b int32
	c [4]byte
	d [4]byte
}
type sliceHead struct {
	p   uintptr
	len int
	cap int
}

var slen = binary.Size(session{})
var sliceSP sliceHead

func testConn() {
	rClient, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 1*1e9, 3*1e9, 2*1e9)
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}
	ret, err := rClient.Do("SET", "name", "zhangsan")
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}
	fmt.Println(ret)
}

const mAXNUM = 100

func main() {
	rClient, err := redis.DialTimeout("tcp4", "127.0.0.1:6379", 3*1e9, 3*1e9, 2*1e9)
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}

	//redis.NewPool()
	v, err := rClient.Do("get", "hello1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1111", v)
	}
	v, err = rClient.Do("setnx", "guangzhou", "23")
	fmt.Println(v, err)
	//getone(rClient)
	//write(rClient)
	//redyLock(rClient)
}

func redyLock(conn redis.Conn) {
	ret, err := conn.Do("set", "ttt", "hello", "nx", "px", 10000)
	if err != nil {
		fmt.Println("1:", err)
	}
	fmt.Println("2", ret)
}

func getone(conn redis.Conn) {
	var mac [6]byte
	ret, err := conn.Do("hget", "device", mac[:])
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("%v\n", ret)
}
func readall() {
	sliceSP.len = 16
	sliceSP.cap = 16
	rClient, err := redis.DialTimeout("tcp4", "127.0.0.1:6379", 3*1e9, 3*1e9, 2*1e9)
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("---------lalalalala----------------%d:%d\n", time.Now().Minute(), time.Now().Second())
	//var mac [8]byte
	//var p = (*int64)(unsafe.Pointer(&mac[0]))
	//for i := 0; i < 1000000; i++ {
	var val = make(map[[6]byte][]byte)
	ret, err := redis.Values(rClient.Do("Hgetall", "device"))
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}
	//	(*p)++
	//}
	//if err=redis.ScanStruct(ret[i],&se);err!=nil{
	//	fmt.Printf("scan error:%v\n", err)
	//        os.Exit(-1)
	//}
	//fmt.Printf("%t================%d\n",ret,len(ret))
	var mac [6]byte
	var se []byte
	for i := 0; i < len(ret); i = i + 2 {
		copy(mac[:], ret[i].([]byte))
		se = ret[i+1].([]byte)
		val[mac] = se
	}
	fmt.Printf("%d\n", len(val))
	//hh:=reflect.ValueOf(ret[1])
	//fmt.Println(hh.Kind(),reflect.TypeOf(ret[1]))
	fmt.Printf("---------lalalalala--------%d:%d\n", time.Now().Minute(), time.Now().Second())
}
func write(conn redis.Conn) {
	var mac [6]byte
	var p = (*int32)(unsafe.Pointer(&mac[0]))
	var se session
	for i := 0; i < mAXNUM; i++ {
		se.a = int32(i)
		se.b = int32(i)
		sliceSP.p = uintptr(unsafe.Pointer(&se))

		//fmt.Printf("====%d====%d\n", slen, binary.Size(session{}))
		_, err := conn.Do("HSET", "hehe", mac[:], *(*[]byte)(unsafe.Pointer(&sliceSP)))
		if err != nil {
			//fmt.Printf("----------set error:%v\n", err)
			os.Exit(-1)
		}
		(*p)++
	}
	//fmt.Printf("---------lalalalala----------------\n")
}

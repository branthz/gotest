package main

import (
	"fmt"
	//"gopkg.in/redis.v2"
	"encoding/binary"
	"github.com/hoisie/redis"
	"os"
	"time"
	"unsafe"
)

var rClient *redis.Client

type session struct {
	a int
	b int
	c [4]byte
	d [4]byte
}

var slen = binary.Size(session{})

func main() {

	rClient = redis.Client(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		PoolSize: 128,
		DB:       0,
	})
	//rClient.Addr = "127.0.0.1:6379"
	//rClient.Db = 13

	write(rClient)

	fmt.Printf("---------lalalalala----------------%d:%d\n", time.Now().Minute(), time.Now().Second())
	//var mac [8]byte
	//var p = (*int64)(unsafe.Pointer(&mac[0]))
	//for i := 0; i < 1000000; i++ {
	var val = make(map[[6]byte]session)
	err := rClient.Hgetall("device", val)
	if err != nil {
		fmt.Printf("----------set error:%v\n", err)
		os.Exit(-1)
	}
	//fmt.Printf("%s\n", val)
	//	(*p)++
	//}
	fmt.Printf("%v\n", val)
	fmt.Printf("---------lalalalala---count:%d-------------%d:%d\n", len(val), time.Now().Minute(), time.Now().Second())
}
func write(conn *redis.Client) {
	var mac [6]byte
	var p = (*int32)(unsafe.Pointer(&mac[0]))
	var se session
	for i := 0; i < 10; i++ {
		se.a = i
		se.b = i

		fmt.Printf("====%d====%d\n", slen, binary.Size(session{}))
		_, err := conn.Hset("device", string(mac[:]), (*(*[]byte)(unsafe.Pointer(&se)))[:slen])
		if err != nil {
			//fmt.Printf("----------set error:%v\n", err)
			os.Exit(-1)
		}
		(*p)++
	}
	fmt.Printf("---------lalalalala----------------\n")
}

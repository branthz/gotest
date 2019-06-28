package main

import (
	"fmt"
	"unsafe"
)

func typeVerify() {
	x := make(map[string]string, 16)
	x["xiaoming"] = "shanghai"
	y := (*struct {
		// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
		// Make sure this stays in sync with the compiler's definition.
		count     int // # live cells == size of map.  Must be first (used by len() builtin)
		flags     uint8
		B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
		noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
		hash0     uint32 // hash seed

		buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
		oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
		nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

		extra uintptr // optional fields
	})(unsafe.Pointer(&x))
	fmt.Printf("%+v\n", y)
}

func typeString() {
	var x = "hello world"
	type stringStruct struct {
		str unsafe.Pointer
		len int
	}
	y := (*stringStruct)(unsafe.Pointer(&x))
	fmt.Printf("%+v\n", y)
}

func typeSlice() {
	var x = make([]int, 10, 20)
	type slice struct {
		array unsafe.Pointer
		ln    int
		cp    int
	}
	y := (*slice)(unsafe.Pointer(&x))
	fmt.Printf("%+v\n", y)
}

/*
func typeChan() {
	x := make(chan int, 100)
	y := (*struct {
		qcount   uint           // total data in the queue
		dataqsiz uint           // size of the circular queue
		buf      unsafe.Pointer // points to an array of dataqsiz elements
		elemsize uint16
		closed   uint32
		elemtype *_type // element type
		sendx    uint   // send index
		recvx    uint   // receive index
		recvq    waitq  // list of recv waiters
		sendq    waitq  // list of send waiters

		// lock protects all fields in hchan, as well as several
		// fields in sudogs blocked on this channel.
		//
		// Do not change another G's status while holding this lock
		// (in particular, do not ready a G), as this can deadlock
		// with stack shrinking.
		lock mutex
	})(unsafe.Pointer(&x))
	fmt.Printf("%+v\n", y)
}*/

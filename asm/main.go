package main

import (
	"fmt"
	"unsafe"
)

type note struct {
	key uintptr
}

type timespec struct {
	tv_sec  int64
	tv_nsec int64
}

const (
	_FUTEX_WAIT = 0
	_FUTEX_WAKE = 1
)

func key32(p *uintptr) *uint32 {
	return (*uint32)(unsafe.Pointer(p))
}

func (ts *timespec) set_sec(x int64) {
	ts.tv_sec = x
}

func (ts *timespec) set_nsec(x int32) {
	ts.tv_nsec = int64(x)
}

//go:noescape
func futex(addr unsafe.Pointer, op int32, val uint32, ts, addr2 unsafe.Pointer, val3 uint32) int32

func noteSleep(n *note, ns int64) {
	var ts timespec

	ts.set_sec(ns / 1000000000)
	ts.set_nsec(int32(ns % 1000000000))
	v := futex(unsafe.Pointer(key32(&n.key)), _FUTEX_WAIT, 0, unsafe.Pointer(&ts), nil, 0)
	fmt.Println(v)
}

func main() {
	no := note{key: 0}
	noteSleep(&no, 1e9)
}

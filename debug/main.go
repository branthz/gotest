package main

import (
	"fmt"
	"runtime/debug"
)

type buffer struct {
	buf [1024]byte
}

func main() {
	var total int = 0
	for i := 0; i < 10; i++ {
		buf := new(buffer)
		total += len(buf.buf)
	}
	debug.PrintStack()
}

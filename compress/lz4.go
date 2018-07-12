package main

import (
	"fmt"

	lz4 "github.com/bkaradzic/go-lz4"
)

func lz4x() {
	//lz4.CompressBound()

	var buf = make([]byte, 2048)
	out, err := lz4.Encode(buf, s1_control_data)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	out2, err := lz4.Encode(buf, out)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("%v===%d\n", out2, len(out2))
}

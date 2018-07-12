package main

import (
	"bytes"
	"fmt"

	lzo "github.com/rasky/go-lzo"
)

func lzox() {
	out := lzo.Compress1X(s1_control_data)
	out2 := lzo.Compress1X(out)
	out3 := lzo.Compress1X(out2)
	fmt.Printf("==%v--------%d,%d\n", out3, len(out3), len(s1_control_data))
	rb := bytes.NewBuffer(out)

	raw, err := lzo.Decompress1X(rb, len(out2), 0)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("======%d\n", len(raw))
}

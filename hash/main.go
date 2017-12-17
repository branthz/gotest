package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

const DEVHASHMAX = 1000000
const STEPLENGTH = 1000000

type entry struct {
	mac   [6]uint8
	using uint16
	next  uint32
}

var entrySize = binary.Size(entry{})
var gEntry []entry

func devHash(mac [6]uint16) int {
	var hash1, hash2 uint32
	hash1 = *(*uint32)(unsafe.Pointer(&mac[0]))
	hash2 = *(*uint32)(unsafe.Pointer(&mac[2]))
	return int(hash1+hash2) & (DEVHASHMAX - 1)
}

func maccmp(mac1, mac2 []byte) bool {
	return bytes.Equal(mac1, mac2)
}

//tp<0 no exist;
func lookUp(mac [6]byte) (int, int) {
	index := devHash(mac)
	for {
		if gEntry[index].using {
			if maccmp(mac, gEntry[index].mac) {
				return index, 1 //using and equal
			} else {
				next := en[index].next
				if next {
					index = next
					continue
				} else {
					return index, -1 //using, no equal,no next,index will be prev node
				}
			}
		} else {
			if gEntry[index].next {
				index = gEntry[index].next
				continue
			} else {
				return index, -2 //no using,no next,index be position perfect
			}
		}
	}
}

func findSpace(mac []byte, pos int) int {
	for {
		if pos < DEVHASHMAX {
			if gEntry[pos+STEPLENGTH].using {

			} else {
				return pos + STEPLENGTH
			}
		} else {

		}
	}
}

func addSession(pen *entry) {
	index, tp := lookUp(devHash(pen.mac))
	if tp == -1 {
		ix = index
		for {
			if gEntry[ix+STEPLENGTH].using {
				ix++
				continue
			}
			break
		}
	}
}

func main() {
	pdata := make([]entry, 2000000)
	for i := DEVHASHMAX; i < STEPLENGTH; i++ {

	}
	var mac [3]uint16
	var pi = (*int32)(unsafe.Pointer(&mac[0]))
	for i := 0; i < 100*10000; i++ {
		hs := devHash(mac)

		if pdata[hs].using {
			fmt.Printf("------------lala bingo!\n")
		}
		pdata[ha].using = 1
		(*pi)++

	}
}

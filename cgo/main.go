package main

import "fmt"

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: ./libhi.so
#include "hi.h" //非标准c头文件，所以用引号
*/
import "C"

func main() {
	C.hi()
	fmt.Println("Hi, vim-go")
}

package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_open(char *name) {
    return shm_open(name, O_RDWR, FILE_MODE);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const SHM_NAME = "my_shm"
const SHM_SIZE = 2 * 1000 * structSize 
const structSize = unsafe.Sizeof(MyData{})

type MyData struct {
	Col1 int32
	Col2 int32
	Col3 int32
}

func main() {
	fd, err := C.my_shm_open(C.CString(SHM_NAME))
	if err != nil {
		fmt.Printf("%v\n",err)
		return
	}

	ptr, err := C.mmap(nil, C.size_t(SHM_SIZE), C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	C.close(fd)

	data := (*[10]MyData)(unsafe.Pointer(ptr))

	fmt.Println(data[0],data[1],data[2])
}

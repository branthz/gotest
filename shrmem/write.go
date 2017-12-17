package main

/*
#cgo linux LDFLAGS: -lrt

#include <fcntl.h>
#include <unistd.h>
#include <sys/mman.h>
#include<sys/types.h>

#define FILE_MODE (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH)

int my_shm_new(char *name) {
    shm_unlink(name);
   	int ret =shm_open(name, O_RDWR|O_CREAT|O_EXCL, 0644);
	if (ret==-1){
		shm_unlink(name);
	}
	return ret;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

const structSize = unsafe.Sizeof(MyData{})
const SHM_NAME = "my_shm"
const SHM_SIZE = 2 * 1000 *1000* structSize

type MyData struct {
	Col1 int32
	Col2 int32
	Col3 int32
	//ha   bool
}

func main() {
	fmt.Printf("-------%d\n", structSize)
	fd := C.my_shm_new(C.CString(SHM_NAME))
	if fd == -1 {
		fmt.Printf("haha:-1\n")
		return
	}

	C.ftruncate(fd, C.__off_t(SHM_SIZE))

	ptr, err := C.mmap(nil, C.size_t(SHM_SIZE), C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	C.close(fd)
	fmt.Println(C.__off_t(SHM_SIZE), C.size_t(SHM_SIZE))
	data := (*[10]MyData)(unsafe.Pointer(ptr))
	fmt.Printf("%p\n", data)
	//data[0].Col1 = 100
	//data[0].Col2 = 876
	//data[1].Col3 = 998
}

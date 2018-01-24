package main

/*
 #include <stdio.h>
 #include <stdlib.h>
 #include "./inc/networkapi_command.h"
 #cgo linux LDFLAGS: -ldnanl -lm -lpthread -L./libs
 #cgo linux CFLAGS: -I../inc
*/
import "C"
import "unsafe"

var path = C.CString("./libs/")

func jsonToBytes(deviceInfo string, controlInfo string, buf []byte) (int, int, int) {
	dstr := C.CString(deviceInfo)
	cstr := C.CString(controlInfo)
	//  path := C.CString("./libs/")
	var cmd C.INT32
	var buflen C.UINT32
	ret := C.cloudserv_json2c(path, dstr, cstr, &cmd, (*C.UINT8)(unsafe.Pointer(&buf[0])), &buflen)
	C.free(unsafe.Pointer(dstr))
	C.free(unsafe.Pointer(cstr))
	return int(ret), int(cmd), int(buflen)
}

func bytesToJson(deviceInfo string, buf []byte, buflen int, cmd int) string {
	dstr := C.CString(deviceInfo)
	//  path := C.CString("./libs/")
	jsonstr := C.cloudserv_c2json(path, dstr, (*C.UINT8)(unsafe.Pointer(&buf[0])), C.UINT32(buflen), C.INT32(cmd))
	C.free(unsafe.Pointer(dstr))
	return C.GoString(jsonstr)
}

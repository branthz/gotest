package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
	"os"
	"time"
)

var s1_control_data = []byte{
	0x5a, 0xa5, 0xaa, 0x55, 0x5a, 0xa5, 0xaa, 0x55,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x09, 0x5f, 0x00, 0x00, 0x22, 0x27, 0xee, 0x03,
	0x94, 0x83, 0x5b, 0x1f, 0x96, 0x0d, 0x43, 0xb4,
	0x01, 0x00, 0x00, 0x00, 0xb2, 0xc7, 0x00, 0x00,
	0x39, 0x40, 0xd0, 0xe6, 0xf3, 0x0b, 0xf6, 0x2d,
	0x9c, 0x7b, 0x56, 0x48, 0xe3, 0x05, 0x6f, 0x75,
	0x37, 0x0d, 0x5c, 0xfb, 0xd9, 0x31, 0xab, 0xc4,
	0xb5, 0x7f, 0x25, 0x8c, 0xa5, 0x37, 0xc0, 0x3f,
	0x77, 0x89, 0xe2, 0xbd, 0xe4, 0x28, 0x4f, 0x4a,
	0xe3, 0xed, 0x30, 0xa7, 0x44, 0x56, 0x8a, 0x4d,
	0x30, 0xaa, 0xf0, 0x89, 0x7c, 0xe8, 0xca, 0x32,
	0x26, 0xd8, 0x83, 0x31, 0x55, 0xe3, 0x44, 0x37,
	0x71, 0x5c, 0x1a, 0x5c, 0x80, 0x22, 0x3b, 0xe6,
	0x45, 0x7f, 0xad, 0xa8, 0x9d, 0x37, 0xfb, 0x03,
	0xee, 0xab, 0x24, 0xf7, 0x1b, 0x06, 0x57, 0xaa,
	0x71, 0xfa, 0x18, 0x24, 0x80, 0x68, 0x7e, 0xa3,
	0xe1, 0xe5, 0x41, 0x35, 0x02, 0xe1, 0x3d, 0x84,
	0xde, 0x00, 0x66, 0xf7, 0xb1, 0xbf, 0xb2, 0x78,
	0x5a, 0xda, 0xcb, 0x29, 0x0c, 0x80, 0x81, 0x2d,
	0x1d, 0xdf, 0x7f, 0x95, 0x56, 0x4f, 0x55, 0xa6,
	0x08, 0xe2, 0x1d, 0x72, 0xb8, 0x20, 0xf4, 0x9a,
	0x2c, 0x9c, 0x76, 0xa1, 0x1f, 0x94, 0x2b, 0x9b,
	0x28, 0x19, 0x8c, 0x6e, 0xc9, 0xa7, 0x0e, 0xb2,
	0xaf, 0xb5, 0xb3, 0xfc, 0xcc, 0x7f, 0x86, 0x68,
	0x23, 0x67, 0xb8, 0x84, 0x33, 0x8b, 0x39, 0xc2,
	0x25, 0x86, 0x7e, 0x96, 0x7a, 0xed, 0xe8, 0x28,
	0x38, 0x04, 0xe9, 0xf4, 0xe5, 0x5b, 0x12, 0x9a,
	0x90, 0x68, 0x92, 0x1b, 0x35, 0x61, 0x0b, 0x88,
	0x91, 0x43, 0x3b, 0x76, 0x92, 0x1c, 0x91, 0xe9,
	0x0c, 0x0e, 0xb0, 0xe8, 0xea, 0xdd, 0xdd, 0xb8,
	0xeb, 0x37, 0x24, 0xc8, 0x5d, 0xb3, 0x8e, 0x3d,
	0xeb, 0xb2, 0xb9, 0xe1, 0xe8, 0x47, 0xac, 0x71,
	0x63, 0x45, 0x13, 0x37, 0x47, 0xdd, 0x39, 0x23,
	0x04, 0xa5, 0x8f, 0x75, 0x40, 0xcb, 0xe1, 0x28,
	0x92, 0x6e, 0xea, 0x02, 0xb9, 0x86, 0x62, 0xe1,
	0x93, 0x8a, 0x6e, 0xa0, 0xa5, 0x45, 0x84, 0xda,
	0xfd, 0x14, 0x3c, 0xa9, 0xb5, 0xd8, 0xcc, 0x58,
	0xf9, 0x45, 0x7a, 0x12, 0xe2, 0x88, 0x28, 0x33,
	0x4a, 0xdb, 0x30, 0x20, 0x0a, 0x08, 0x87, 0x71,
	0x91, 0xe3, 0x3f, 0x37, 0xfe, 0x32, 0xbc, 0x77,
	0xb0, 0x85, 0xe0, 0xbb, 0x34, 0xb2, 0x3f, 0x29,
	0x08, 0xdd, 0x04, 0xd7, 0xd7, 0x21, 0xc2, 0x3d,
	0x78, 0xd8, 0xb6, 0x36, 0xdf, 0xb9, 0xa9, 0x37,
	0x6c, 0xe4, 0x7e, 0xd0, 0xb6, 0x72, 0x5b, 0xcc,
	0x1c, 0x18, 0x41, 0xe8, 0x48, 0x4a, 0x16, 0x75,
	0xa7, 0xfd, 0xd5, 0xf7, 0xe9, 0x3a, 0x9b, 0x74,
	0x0a, 0xb2, 0xff, 0xb4, 0x28, 0xe5, 0x81, 0x75,
	0x65, 0x09, 0x2f, 0xb5, 0xc5, 0xbd, 0x12, 0xfa,
	0xe7, 0x85, 0x8e, 0x8e, 0x67, 0x28, 0x88, 0x2f,
	0x0f, 0x88, 0xab, 0x83, 0x18, 0x0f, 0x88, 0x7c,
	0x4d, 0x46, 0x67, 0x41, 0x50, 0xc7, 0xc2, 0x4a,
	0x7b, 0x18, 0x4c, 0x19, 0xdd, 0x1b, 0xc5, 0xd9,
	0xf3, 0x1d, 0x2f, 0x54, 0xa2, 0xa0, 0x0f, 0x88,
	0x54, 0x68, 0xac, 0x1e, 0x1e, 0xd5, 0x70, 0x03,
	0xdb, 0x46, 0xe3, 0xa8, 0xb1, 0x0c, 0x13, 0xd3,
	0xbe, 0x02, 0xb2, 0x64, 0x26, 0x49, 0xce, 0x1f,
	0xe2, 0x6e, 0xdc, 0x8b, 0x8d, 0x9f, 0x7d, 0xd8,
	0x98, 0xb2, 0xde, 0x20, 0x03, 0x1a, 0xe2, 0xdb,
	0xcd, 0x84, 0x73, 0xd8, 0x2d, 0x25, 0x98, 0xa6,
	0x14, 0xce, 0x7e, 0x89, 0xa9, 0xe6, 0x55, 0x59,
	0x28, 0xb7, 0x04, 0xb8, 0x83, 0xfd, 0xe6, 0x1f,
	0xe2, 0xaa, 0x3e, 0x18, 0xf4, 0x52, 0x37, 0x39,
	0x61, 0x03, 0xf2, 0x02, 0x19, 0x0f, 0x88, 0x7f,
	0x71, 0xe4, 0xc9, 0x86, 0x33, 0x8f, 0x00, 0x88,
	0x0f, 0x1b, 0xba, 0x44, 0xcf, 0xef, 0x59, 0xeb,
	0x91, 0x8e, 0x4f, 0xcb, 0xea, 0x8e, 0x87, 0x56,
	0x82, 0x7a, 0xae, 0x70, 0x5c, 0xd9, 0x47, 0x4b,
	0x08, 0x64, 0xbd, 0x7f, 0x3e, 0x01, 0x3b, 0xd7,
	0xe8, 0xc0, 0xa9, 0x7c, 0x6d, 0x9e, 0x19, 0x77,
	0x68, 0xbd, 0x7c, 0x37, 0x79, 0xcb, 0x12, 0x72,
	0xd2, 0xd8, 0x7b, 0x73, 0xc6, 0x8d, 0xe8, 0x0a,
	0x6c, 0x8d, 0x64, 0x62, 0xdb, 0x92, 0xd4, 0xdd,
	0xf5, 0x51, 0x93, 0x0c, 0x30, 0xa5, 0x5c, 0xe3,
	0xff, 0xbd, 0x2b, 0xdd, 0xc2, 0xd2, 0x6c, 0xf5,
	0x0b, 0x1c, 0x23, 0xbe, 0xfd, 0xd7, 0xea, 0x2f,
	0xfe, 0xd6, 0x40, 0xec, 0xdf, 0x63, 0x48, 0xe7,
	0x9c, 0x9f, 0xa0, 0xa0, 0x4a, 0xf0, 0xc9, 0x06,
	0x58, 0x76, 0x9e, 0xaf, 0x80, 0xab, 0xb6, 0x2f,
	0x35, 0x95, 0x28, 0x20, 0x9e, 0x77, 0xf4, 0x97,
	0xd7, 0x4e, 0x0e, 0xc3, 0x20, 0xaa, 0x56, 0x69,
	0x20, 0x0f, 0x2b, 0x61, 0x89, 0x7a, 0x61, 0x7c,
	0xe9, 0x39, 0xd2, 0x55, 0xaa, 0x0c, 0x3e, 0x4b,
	0xd5, 0x87, 0x7b, 0x49, 0x0a, 0x84, 0xfa, 0x84,
	0x25, 0x0f, 0x7f, 0x87, 0x56, 0x0b, 0x45, 0x80,
	0x88, 0x94, 0x7e, 0x97, 0x3e, 0xca, 0x22, 0xbc,
	0xe5, 0x43, 0x28, 0xc2, 0x5b, 0xa0, 0x6c, 0x02,
	0x7c, 0x82, 0x36, 0x99, 0x0f, 0x9d, 0x76, 0xce,
	0x64, 0x7d, 0xe9, 0x90, 0xf4, 0x9f, 0xf4, 0x5b,
	0x95, 0xbe, 0xa3, 0xc1, 0xf1, 0xe9, 0xa0, 0xa1,
	0x48, 0x26, 0xd0, 0x2e, 0x45, 0x9e, 0xe2, 0x83,
	0x1e, 0x0e, 0x2c, 0xce, 0xbb, 0x6f, 0x53, 0x16,
	0x99, 0x8d, 0xd7, 0x0a, 0x8d, 0x9e, 0x32, 0x7b,
	0xb0, 0xa6, 0x41, 0xfb, 0x4a, 0x81, 0x27, 0x04,
	0x9e, 0x0f, 0x3c, 0xef, 0x8f, 0xdc, 0x0a, 0xa9,
	0xe1, 0x3a, 0x56, 0xf7, 0xa1, 0x4a, 0xbd, 0x49,
	0xf8, 0x5e, 0x4b, 0xe8, 0xac, 0x46, 0x7a, 0x92,
	0xc0, 0x42, 0x4c, 0x87, 0x63, 0x53, 0x8e, 0x81,
	0xd9, 0x5d, 0x89, 0xe2, 0xaa, 0xf8, 0x50, 0x60,
	0xe4, 0x7b, 0x42, 0x38, 0xef, 0x1b, 0xf4, 0x1c,
	0xe3, 0x95, 0xa8, 0x09, 0x22, 0xe3, 0x51, 0xb3,
	0xae, 0x04, 0xce, 0x6b, 0xdf, 0x37, 0x1e, 0xe4,
	0x3a, 0xbd, 0x84, 0x3e, 0x85, 0x60, 0x80, 0x87,
	0xa4, 0x84, 0x0a, 0xf3, 0xf6, 0x2f, 0x46, 0xc3,
	0x22, 0x6b, 0x2e, 0xa5, 0xf7, 0xc2, 0x9a, 0x69,
	0x26, 0x5d, 0x48, 0xc0, 0x4a, 0x16, 0xc0, 0x4e,
	0x9f, 0x22, 0x22, 0x64, 0xcf, 0xdb, 0x72, 0xce,
	0x61, 0xed, 0x2d, 0x66, 0x98, 0xae, 0x9b, 0x64,
	0x4f, 0x9a, 0x4c, 0xcd, 0x3b, 0xc7, 0xdd, 0x96,
	0x4a, 0xff, 0x58, 0x16, 0x59, 0x94, 0x08, 0x2e,
	0xdd, 0x66, 0xe6, 0x2c, 0xea, 0xc2, 0xf2, 0x34,
	0xf3, 0x0b, 0x9c, 0x8c, 0x93, 0x5a, 0x86, 0x29,
	0xbf, 0x9b, 0x7e, 0xed, 0x31, 0x84, 0x17, 0x82,
	0xc5, 0x8d, 0x83, 0x31, 0x7d, 0x99, 0x6e, 0x1c,
	0x04, 0x02, 0x7c, 0x5f, 0x45, 0x40, 0xe8, 0x1d,
	0x75, 0x74, 0xfb, 0x09, 0x50, 0x45, 0x88, 0x1d,
	0xed, 0x26, 0xff, 0xc3, 0x1b, 0x60, 0xd7, 0x45,
	0xdb, 0x33, 0xe2, 0x34, 0xff, 0x34, 0x57, 0xa1,
	0x93, 0x4c, 0x9b, 0xcb, 0x30, 0xec, 0x2a, 0xe7,
	0x97, 0x43, 0x2e, 0xff, 0xe5, 0xfd, 0x42, 0xd9,
	0x75, 0xa5, 0x9e, 0x4a, 0x54, 0x22, 0x55, 0x28,
	0x75, 0xb1, 0x3f, 0x4a, 0xe8, 0xcd, 0xbb, 0x5b,
	0xe6, 0x70, 0x1a, 0x30, 0x74, 0x20, 0x82, 0xb6,
	0x68, 0x55, 0x91, 0x68, 0x31, 0x6c, 0x17, 0x9a,
	0x72, 0x76, 0x0e, 0x78, 0x1d, 0x68, 0x5b, 0x68,
	0x0e, 0xcf, 0x6f, 0x8e, 0x65, 0x90, 0x0c, 0x18,
	0xd5, 0x13, 0x63, 0x7c, 0x1a, 0xae, 0xf9, 0xff,
	0x89, 0xc2, 0xce, 0xb8, 0xa2, 0x61, 0x74, 0xab,
	0x74, 0x8b, 0x29, 0xa3, 0x7a, 0xb1, 0xc9, 0x88,
	0x57, 0x80, 0x42, 0x4b, 0xe1, 0xf1, 0x82, 0x25,
	0x2e, 0x88, 0x33, 0x7b, 0xaf, 0xa6, 0x6b, 0xc1,
	0x38, 0xab, 0x55, 0x5e, 0xf9, 0x73, 0xf8, 0xc2,
	0x1a, 0xa3, 0x09, 0x56, 0x80, 0x77, 0x8e, 0x5d,
	0x8d, 0x3b, 0xfc, 0xb4, 0xb9, 0x73, 0x9d, 0x96,
	0x0a, 0xd4, 0x71, 0x67, 0x1d, 0xed, 0x2d, 0xeb,
	0x03, 0xb8, 0x7a, 0x79, 0xfa, 0xce, 0xe0, 0x15,
	0xb5, 0x56, 0xe4, 0x69, 0xcf, 0xfc, 0x9b, 0xb5,
	0x54, 0xd2, 0xd2, 0xc1, 0xd8, 0xb2, 0x02, 0xf8,
	0x68, 0x9e, 0x06, 0xb4, 0x96, 0x24, 0x5a, 0x58,
	0xb9, 0x9f, 0x0a, 0xce, 0xbb, 0xff, 0xaf, 0x7e,
	0x26, 0xe7, 0x42, 0x9e, 0x03, 0x68, 0x49, 0x17,
	0x37, 0xe0, 0xe0, 0x65, 0x27, 0x27, 0xe6, 0x9a,
	0x85, 0x9f, 0x87, 0x0e, 0x3f, 0x95, 0x1a, 0x04,
	0x7b, 0xeb, 0x3a, 0xd6, 0x18, 0x51, 0x5e, 0x67,
	0x09, 0x62, 0x85, 0x62, 0x68, 0xb8, 0x0e, 0x0e,
	0x57, 0x64, 0x08, 0x24, 0x84, 0x87, 0xed, 0x03,
	0xbb, 0xfc, 0x7b, 0x8b, 0xab, 0x78, 0xa4, 0x40,
	0xfb, 0x5f, 0xd4, 0x34, 0x68, 0x82, 0x79, 0x23,
	0xa9, 0x5f, 0x16, 0x0f, 0x93, 0x3e, 0x99, 0xfd,
	0x69, 0x19, 0x62, 0x5e, 0xdf, 0x53, 0xae, 0xc4,
	0x60, 0xab, 0xe5, 0xbf, 0xc1, 0x71, 0xdb, 0x75,
	0x5c, 0x4b, 0xb7, 0xe4, 0xe7, 0x3e, 0x4b, 0x78,
	0xc2, 0xb2, 0xae, 0xe4, 0xdb, 0x67, 0xf3, 0x34,
	0xbf, 0x3d, 0x45, 0x05, 0xcc, 0x2e, 0xcf, 0xc9,
	0xcf, 0x17, 0x14, 0x9d, 0x56, 0x2e, 0x7a, 0xd9,
	0x88, 0x18, 0x77, 0x76, 0x00, 0x21, 0x3f, 0x5b,
	0x70, 0x16, 0x84, 0x30, 0x50, 0xb5, 0xfc, 0x8b,
	0xdb, 0x9e, 0x51, 0x15, 0xeb, 0xf0, 0x0c, 0x9a,
	0x5e, 0xd7, 0x05, 0x42, 0xc9, 0x34, 0x1f, 0xb5,
	0x05, 0xcf, 0xee, 0x54, 0x0c, 0x9a, 0x6d, 0x42,
	0x39, 0x00, 0x31, 0x06, 0x7b, 0xcb, 0x2e, 0x36,
	0xde, 0xdd, 0x91, 0x3b, 0x1f, 0xc6, 0x7d, 0xd9,
	0x42, 0x31, 0x02, 0x69, 0xcf, 0xdc, 0x5e, 0x07,
	0xf9, 0xd3, 0x0d, 0x1f, 0xee, 0x4d, 0x54, 0x83,
	0x1f, 0x5a, 0x50, 0x18, 0x21, 0x3a, 0x78, 0xd5,
	0xf3, 0x40, 0x24, 0xc7, 0x35, 0x8e, 0x08, 0x05,
	0x0d, 0x19, 0x68, 0x70, 0x13, 0x02, 0x7e, 0x4c,
	0xd3, 0xa5, 0xd6, 0x75, 0x0a, 0x69, 0xb8, 0x23,
	0x0f, 0x11, 0x41, 0x7a, 0xc5, 0xba, 0xe6, 0x94,
	0x0f, 0xbe, 0x2d, 0x05, 0x2a, 0xa7, 0xc3, 0x50,
	0x2f, 0x91, 0x9e, 0x50, 0xde, 0x74, 0xdf, 0x37,
	0x57, 0xce, 0x88, 0x18, 0x54, 0x95, 0x07, 0x89,
	0x65, 0x49, 0xf2, 0xc1, 0xfd, 0x61, 0xaa, 0x20,
	0x6b, 0x88, 0xb9, 0xb8, 0xba, 0x46, 0x5b, 0x57,
	0x83, 0x65, 0x78, 0x2a, 0x2c, 0x52, 0xd9, 0xd1}

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
}

func validFlate() {
	log.Printf("raw length:%d\n", len(s1_control_data))
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, -1)
	if err != nil {
		Logger.Fatal(err)
		return
	}

	w.Write(s1_control_data)
	w.Flush()
	Logger.Printf("compressed data first time:%v----length:%d\n", b.Bytes(), b.Len())

	var tbuf = make([]byte, 2048)
	tbuflen := copy(tbuf, b.Bytes())

	w.Reset(&b)
	w.Write(tbuf[:tbuflen])

	w.Close()

	Logger.Printf("compressed data:%v----length:%d\n", b.Bytes(), b.Len())

	//br := bytes.NewReader(b.Bytes())

	fr := flate.NewReader(bytes.NewReader(b.Bytes()))
	var data = make([]byte, 4096)
	rn, err := fr.Read(data)
	if err != nil {
		Logger.Fatal(err)
		return
	}
	fr = flate.NewReader(bytes.NewReader(data[:rn]))
	var xxx = make([]byte, 4096)
	rn, err = fr.Read(xxx)
	if err != nil {
		Logger.Fatal(err)
		return
	}
	Logger.Printf("decompressed data:%v------length:%d\n", xxx[:rn], rn)
	fr.Close()
}

func benchmark() {
	var src = make([]byte, 1500)
	for i := 0; i < len(src); i++ {
		src[i] = byte(i)
	}
	for i := 500; i < 1000; i++ {
		src[i] = 0
	}
	var tm = time.Now().Unix()
	for i := 0; i < 10000; i++ {

		var b bytes.Buffer
		w, err := flate.NewWriter(&b, -1)
		if err != nil {
			fmt.Println("+++++++++++++++++", err)
			return
		}

		w.Write(src)
		w.Close()
		/*
			r, err := zlib.NewReader(&b)
			if err != nil {
				fmt.Println("+++++++++++++++++", err)
				return
			}
			n, _ := io.Copy(os.Stdout, r)
			fmt.Printf("%d------\n", n)
			r.Close()
		*/

		//fmt.Printf("compressed data:%v----length:%d\n", b.Bytes(), b.Len())

		br := bytes.NewReader(b.Bytes())

		fr := flate.NewReader(br)
		var data = make([]byte, 2048)
		_, err = fr.Read(data)
		if err != nil {
			fmt.Println("+++++++++++++++++", err)
			return
		}
		//fmt.Printf("decompressed data:%v------length:%d\n", data[:rn], rn)
		fr.Close()
	}
	fmt.Printf("hahah:%d\n", time.Now().Unix()-tm)
}

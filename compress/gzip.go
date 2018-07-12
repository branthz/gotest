package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

// Write gzipped data to a Writer
func gzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	gw.Flush()
	return err
}

// Write gunzipped data to a Writer
func gunzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

func CheckGzip() {
	var st = `{ "version": "v1.3.1", "timestamp_seconds": "17:15:04"}`
	RawData = []byte(st)
	fmt.Println("original:\t", len(RawData))

	var buf bytes.Buffer
	err := gzipWrite(&buf, RawData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("compressed:\t", len(buf.Bytes()), buf.Bytes())

	var buf2 bytes.Buffer
	err = gunzipWrite(&buf2, buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	graw := buf2.String()
	if graw != string(RawData) {
		fmt.Printf("not correct\n")
	} else {
		fmt.Println("ok!")
	}
}

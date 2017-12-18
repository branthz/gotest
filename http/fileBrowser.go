package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"net/url"
	//"errors"
	"package/tools"
	//"sync"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/1", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("/home/brant/temp")))
	http.ListenAndServe(":9090", nil)
}


func uploadHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/home/brant/Documents/kit接口.docx")
	if err != nil {
		fmt.Printf("failed create file:%v", err)
		return
	}
	defer f.Close()
	body1, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	w.Write(body1)
	//fmt.Fprintf(w, string(body1))
	return
}



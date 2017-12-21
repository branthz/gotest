package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/grace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WORLD++++=222228989!"))
	w.Write([]byte("ospid:" + strconv.Itoa(os.Getpid())))
	time.Sleep(1e9 * 5)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)

	err := grace.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 8080 stopped")
	os.Exit(0)
}

package main

import(
 	_ "net/http/pprof"
	"net/http"
)


//浏览器访问http://localhost:6060/debug/pprof/
func main(){
	http.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){
		w.Write([]byte("server is ok"))
	})
	http.ListenAndServe("localhost:6060", nil)
}


package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}
func main() {
	http.Handle("/foo", new(countHandler))
	http.ListenAndServe(":5555", nil)
}

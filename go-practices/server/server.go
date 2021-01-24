package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Millisecond)
	fmt.Fprintf(w, "hello\n")
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

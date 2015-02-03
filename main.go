package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello, world</h1>\nI'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
  fmt.Println("running")
}

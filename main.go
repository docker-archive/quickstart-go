package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/mgo.v2"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello, world</h1>\nI'm running on %s with an %s CPU the name of this program is : %s", runtime.GOOS, runtime.GOARCH, os.Getenv("NAME"))
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO"))
	if err != nil {
		fmt.Println("No mongo db")
	}
	defer session.Close()
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

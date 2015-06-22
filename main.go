package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

func mongoConnect() (s string) {

	session, err := mgo.Dial("mongo")
	if err == nil {
		defer session.Close()
		s = "connected"
	} else {
		s = "not available"
	}
	return s
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	hostname, _ := os.Hostname()
	mongostatus := mongoConnect()
	fmt.Fprintf(w, "<h1>hello, %s</h1>\n<b>Hostname: </b>%s<br><b>MongoDB Status: </b>%s", os.Getenv("NAME"), hostname, mongostatus)
	fmt.Println(hostname, "handled HTTP REQUEST at", time.Now(), "\nMongoDB Status:", mongostatus)
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Everything is working !\n Listening on port 80 for requests...")
	http.ListenAndServe(":80", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Type string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "<h1>hello, %s</h1>\n<b>Hostname:</b> %s<br/>", os.Getenv("NAME"), hostname)
}

func main() {
	session, err := mgo.Dial("mongo")
	if err == nil {
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("quicksell").C("people")
		err = c.Insert(&Person{"Golang", "Awesome"})
		if err != nil {
			log.Fatal(err)
		}
		result := Person{}

		err = c.Find(bson.M{"name": "Golang"}).One(&result)
		fmt.Println("Golang is:", result.Type)
	} else {
		fmt.Println("Please link a service named \"mongo\" to this service.")
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

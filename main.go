package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Type string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello, world</h1>\nI'm running on %s with an %s CPU.\n The name of this program is : %s and runs with a MongoDB database named %s", runtime.GOOS, runtime.GOARCH, os.Getenv("NAME"), os.Getenv("MONGO_URL"))
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
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
		fmt.Println("Golang is: ", result.Type)
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

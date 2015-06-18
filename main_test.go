package main

import "testing"

func Test_mongoConnect(t *testing.T) {
	s := mongoConnect()
	if s == "Not available" {
		t.Error("MongoDB is currently " + s)
	}
}

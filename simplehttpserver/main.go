package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("/tmp/"))  // file server root
	http.Handle("/", fs) // register handle

	http.ListenAndServe(":80", nil)
}
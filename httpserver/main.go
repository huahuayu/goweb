package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("/Users/shiming/Nutstore/5-goland/src/goweb/httpserver/static"))  // file server root
	http.Handle("/static/", http.StripPrefix("/static/", fs)) // StripPrefix usage refer https://stackoverflow.com/a/47997908/6797425

	http.ListenAndServe(":80", nil)
}
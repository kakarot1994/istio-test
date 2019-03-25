package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, " My World!")
}

func main() {
	http.HandleFunc("/", hello1)
	http.ListenAndServe(":80", nil)
}

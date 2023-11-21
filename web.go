package main

import (
	"fmt"
	"net/http"
)

var port = ":9098"

func main() {
	http.HandleFunc("/", greet)

	fmt.Println("Web Started ", port)
	http.ListenAndServe(port, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("service '/' ", "success")
	msg := "hello world"
	fmt.Fprint(w, msg)
}

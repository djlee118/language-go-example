package main

import (
	"fmt"
	"log"
	"net/http"
)

func magic_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, magic world.")
}

func main() {
	http.HandleFunc("/", magic_handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

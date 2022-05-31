package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service")

	http.HandleFunc("/", BaseHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from account service")
}

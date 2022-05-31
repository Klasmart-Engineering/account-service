package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service")

	http.HandleFunc("/", RootHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from account service")
}

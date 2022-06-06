package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting KidsLoop Account Service on http://localhost:8080")

	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8080", nil)
}

func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from account service")
}

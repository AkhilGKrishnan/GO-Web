package main

import (
	"fmt"
	"net/http"
)

// HandlerFunc function
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wel to the world of hacking")
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	http.ListenAndServe(":3010", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port :9000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.ListenAndServe(":9000", nil)
}

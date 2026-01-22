package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Jenkins!")
	})
	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Service-demo!")
	})
	http.HandleFunc("/harbor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Harbor!")
	})
	http.ListenAndServe(":80", nil)
}

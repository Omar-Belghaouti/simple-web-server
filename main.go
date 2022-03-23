package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if r.URL.Path != "/form" {
		http.NotFound(w, r)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	fmt.Fprintf(w, "Form data:\n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "  %q: %q\n", k, v)
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server on port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

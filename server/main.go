package main

import (
	"fmt"
	"net/http"
)

func hander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Server")
}
func homehander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello i am homepage")
}
func main() {
	http.HandleFunc("/", hander)
	http.HandleFunc("about", homehander)
	http.ListenAndServe("8080", nil)
}

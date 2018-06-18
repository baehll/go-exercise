package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/liss", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HTTP PARAMETER")
	fmt.Fprintf(w, "r.Form: %q\n", r.Form["cycles"])
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func lissajousHandler(w http.ResponseWriter, r *http.Request){
	lissajous(w)
}
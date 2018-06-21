package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
)

type Page struct {
	Title string
	body []byte
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r* http.Request) {
	title := r.URL.Path[1:]
	p := loadView(title)
	body := string(p.body[:])
	fmt.Fprintf(w, "%q", body)	
}

func loadView(title string) *Page {
	if (strings.HasSuffix(title, ""))
	filename := title + ".txt"
	b, err := ioutil.ReadFile("./" + filename)
	log.Print("./" + filename)
	if err != nil {
		return &Page{Title: "Error", body: []byte(err.Error())}
	} else {
		return &Page{Title: title, body: b}
	}
}
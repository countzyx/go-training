package main

import (
	"fmt"
	"net/http"
)

type Default struct{}

func (d Default) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "<html><head><title>HTTP Handlers</title></head><body><a href='/string'>String</a><br/><a href='/struct'>Struct</a></body></html>")
}

type String string

func (s String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, string(s))
}

type Struct struct {
	Greeting, Punct, Who string
}

func (s *Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, " ", s.Who)
}

func main() {
	var d Default
	http.Handle("/", d)
	http.Handle("/string", String("I'm a frayed know."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}

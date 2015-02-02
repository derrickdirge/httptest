package main

import (
	"log"
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct	 string
	Who	 string
}

func (s String) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s Struct) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/string", String("Abbey is the bomb"))
	http.Handle("/struct", &Struct{"My man", ",", "Gusmo!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

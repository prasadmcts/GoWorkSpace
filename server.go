package main

import (
	"fmt"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)
	fmt.Println("http://localhost:8080/SomeName")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

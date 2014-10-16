package main

import (
	"fmt"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:]
	fmt.Fprintf(w, "Hi there, I love %s!", s)
}
 

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}


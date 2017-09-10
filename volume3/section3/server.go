package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)

}

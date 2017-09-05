package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	fs := http.FileServer(http.Dir("./static"))
	flag.Parse()
	log.SetFlags(0)
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

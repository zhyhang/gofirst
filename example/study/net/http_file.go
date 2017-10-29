package main

import (
	"net/http"
	"log"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/data/webdir")))
	http.Handle("/base/", http.StripPrefix("/base/", http.FileServer(http.Dir("/data/webdir/test"))))
	log.Fatal(http.ListenAndServe(":8080",nil))
}

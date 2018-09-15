package server

import (
	"log"
	"net/http"
)

func FileServer() {
	http.Handle("/", http.FileServer(http.Dir("/data/webdir")))
	http.Handle("/base/", http.StripPrefix("/base/", http.FileServer(http.Dir("/data/webdir/test"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

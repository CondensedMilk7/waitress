// A simple little static file server.
// Can be used to serve custom browser start page.

package main

import (
	"flag"
	"log"
	"net/http"
)

var path string
var port string

func main() {

	flag.StringVar(&path, "path", "", "Serve directory")
	flag.StringVar(&port, "port", "8080", "Server port")
	flag.Parse()

	if path == "" {
		log.Fatal("A path must be provided!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path+r.URL.Path)
	})

	log.Printf("Serving %s on port %s", path, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

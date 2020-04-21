package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var CommitSHA string

var (
	wwwRoot = flag.String("root", "/usr/share/booktaxi", "path to serve files from")
	port    = flag.Int("port", 8080, "port to listen on")
)

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*wwwRoot))
	http.Handle("/", addVersionHeader(fs))

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("version %s listening on %s\n", CommitSHA[:6], addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func addVersionHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("DemoVersion", CommitSHA)
		h.ServeHTTP(w, r)
	})
}

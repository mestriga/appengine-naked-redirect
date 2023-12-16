package main

import (
	"log"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=86400")
	w.Header().Add(
		"Strict-Transport-Security",
		"max-age=31536000; includeSubDomains; preload")
	rurl := r.URL
	rurl.Scheme = "https"
	rurl.Host = "www." + r.Host
	http.Redirect(w, r, rurl.String(), http.StatusTemporaryRedirect)
}

func main() {
	if err := http.ListenAndServe(
		":" + os.Getenv("PORT"),
		http.HandlerFunc(handleRequest));
		err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("page/")))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

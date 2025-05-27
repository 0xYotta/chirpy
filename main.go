package main

import (
	"log"
	"net/http"
)

func main() {
	// consts
	const port = "8080"
	const filepathRoot = "."

	// mux & server
	mux := http.NewServeMux() // I: returning pointer so no need to make it manually

	server := &http.Server{
		// NOTE: using & because structure is huge.
		// it's idiomatic to use pointers in this cases

		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// logs
	log.Printf("Serving on port: %v", port)
	log.Fatal(server.ListenAndServe())
}

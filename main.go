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

	mux.HandleFunc("/healthz", HandleHealthz)
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	// logs
	log.Printf("Serving on port: %v", port)
	log.Fatal(server.ListenAndServe())
}

func HandleHealthz(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/plain; charset=utf-8")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte("OK"))
}

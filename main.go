package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("ready to go!")

	mux := http.NewServeMux()

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}

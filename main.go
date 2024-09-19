package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func printRequestDetail(r *http.Request) {
	log.Printf("Request from %s: %s %s", r.RemoteAddr, r.Method, r.URL)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		printRequestDetail(r)
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		printRequestDetail(r)
		fmt.Fprintf(w, "I'm healthy!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

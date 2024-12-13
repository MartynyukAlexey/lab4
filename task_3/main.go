package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	staticDir := "./"

	http.Handle("/", http.FileServer(http.Dir(staticDir)))

	log.Printf("Starting server on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

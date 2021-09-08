package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Print something on startup
	fmt.Println("Starting HTTP server!")
	port := os.Getenv("PORT")
	fmt.Println("Running on PORT::", port)

	// Routes
	// Print msg "Hello, world!" or "Hello, URL!"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// Start server
	http.ListenAndServe(":"+port, nil)
}

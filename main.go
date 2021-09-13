package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ForestFlowerr/portfolio/app/controllers"
)

func main() {

	// Print something on startup
	fmt.Println("Starting HTTP server!")
	port := os.Getenv("PORT")
	fmt.Println("Running on PORT::", port)

	hello := controllers.Hello

	// Routes
	// Print msg "Hello, world!" or "Hello, URL!"
	http.HandleFunc("/", hello)

	// Start server
	http.ListenAndServe(":"+port, nil)
}

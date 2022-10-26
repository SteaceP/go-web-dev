package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steacep/hello-world/v2/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, nil))
}

package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const portNumber = ":8081"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the home page!\n")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.00, 10.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Printf("%f divided by %f is %f\n", 100.0, 10.0, f)
}

// divideValues divide two float32 and return an error if divided by 0
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s\n", portNumber)

	log.Fatal(http.ListenAndServe(portNumber, nil))
}

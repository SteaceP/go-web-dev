package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		n, err := io.WriteString(w, "Hello, world!\n")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Number of bytes written: %d \n", n)
	}

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

//? doesn't work for some reason, I took the example of go.dev
//?  and modified it to have the same output
// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		n, err := fmt.Fprintf(w, "Hello World")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("Number of bytes written: %d \n", n)
// 	})

// 	_ = http.ListenAndServe(":8080", nil)
// }

package main

import "log"

//? Function with 2 return value
func main() {
	var whatToSay string
	var saySomethingElse string

	whatToSay, _ = saySomething("Hello!") // underscore is used to ignore the second return value
	log.Println(whatToSay)

	saySomethingElse, _ = saySomething("Goodbye")
	log.Println(saySomethingElse)

	log.Println(saySomething("Finally"))
}

func saySomething(s string) (string, string) {
	return s, "World"
}

//? Function with a single return value and variables basics
// func main() {
// 	var whatToSay string
// 	var saySomethingElse string
// 	var i int

// 	whatToSay = saySomething("Hello!")
// 	log.Println(whatToSay)

// 	saySomethingElse = saySomething("Goodbye")
// 	log.Println(saySomethingElse)

// 	log.Println(saySomething("Finally"))

// 	i = 7 // Println -> 7
// 	i = 8 // Println -> 8
// 	log.Println(i) // -> 0   (if i is not defined explicitly)
// }

// func saySomething(s string) string {
// 	return s
// }

package main

import "log"

// ? 'make()` is used to initializes an object(map, slice, chan)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)
}

// ? Slice

//* initialize a Slice
//! var mySlice []int

//* append `int 2` to mySlice
//! mySlice = append(mySlice, 2)

//* sort (in order) mySlice
//! sort.Ints(mySlice) // -> [2]

//* initialise a Slice with predefined values and assign it to a variable
//! numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

//! log.Println(numbers) // -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
//! log.Println(numbers[6:9]) // -> [7, 8, 9]

//

// ? Map

// ? maps can old any value of any type and as much as you need
// ? not necessarily in order, always access them via the Key
// ? it's mostly like object in JS

//* initialize a map that will have the `Key` as a string and will return the User type
//! myMap := make(map[string]User)

//* set some data from the User type
//!	me := User{
//!		FirstName: "Steacy",
//!		LastName:  "Paquette",
//!	}

//* map the previous data
//!	myMap["me"] = me

//* a simple way to use it
//! log.Println(myMap["me"].FirstName, myMap["me"].LastName)

package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var mySlice []User

	u1 := User{
		FirstName: "Steace",
	}

	u2 := User{
		FirstName: "Sam",
	}

	mySlice = append(mySlice, u1)
	mySlice = append(mySlice, u2)

	for i, x := range mySlice {
		fmt.Println(i, x.FirstName)
	}

}

//* iterate over a map
// myMap := make(map[string]string)
// myMap["dog"] = "dog"
// myMap["fish"] = "fish"
// myMap["cat"] = "cat"

// for i, x := range myMap {
// 	fmt.Println(i, x)
// }

//* iterate over a slice
// mySlice := []string{"dog", "cat", "horse", "fish", "banana"}

// for _, x := range mySlice {
// 	fmt.Println(x)
// }

//* Basic for loop
// for i := 0; i <= 10; i++ {
// 	fmt.Println(i)
// }

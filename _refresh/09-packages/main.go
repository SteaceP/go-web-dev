package main

import (
	"fmt"

	"github.com/steacep/testprogram/helpers"
)

func main() {
	fmt.Println("hello")

	var myVar helpers.SomeType
	fmt.Println(myVar)

}

//? go mod init github.com/steacep/{name-of-package}

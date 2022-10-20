package main

import (
	"fmt"

	"github.com/steacep/somepackage/helpers"
)

const numPool = 9000000000000000000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}

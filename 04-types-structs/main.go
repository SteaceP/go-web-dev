package main

import (
	"log"
	"time"
)

var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Steace",
		LastName:    "Paquette",
		PhoneNumber: "555-555-5555",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}

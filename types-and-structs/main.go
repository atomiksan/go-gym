package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Satya",
		LastName:  "Sena",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate :", user.BirthDate)
}

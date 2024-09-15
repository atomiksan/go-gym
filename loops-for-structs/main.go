package main

import "log"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func main() {
	var users []User

	users = append(users, User{"John", "Smith", "john@smith.com", 28})
	users = append(users, User{"Mary", "Brown", "mary@brown.com", 25})
	users = append(users, User{"Sally", "Watson", "sally@watson.com", 27})
	users = append(users, User{"Alex", "Jones", "alex@Jones.com", 32})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}

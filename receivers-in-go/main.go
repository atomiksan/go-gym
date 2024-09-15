package main

import "log"

type myStruct struct {
	FirstName string
	LastName  string
}

func (m *myStruct) printFirstname() string {
	return m.FirstName
}

func main() {
	var v1 myStruct
	v1.FirstName = "Satya"

	v2 := myStruct{
		FirstName: "Sena",
	}

	log.Println("v1's First Name is:", v1.FirstName)
	log.Println("v2's First Name is:", v2.printFirstname())
}

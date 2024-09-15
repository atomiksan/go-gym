package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Satya",
		LastName:  "Sena",
	}
	myMap["me"] = me

	log.Println(myMap["me"])

	var mySlice []string

	mySlice = append(mySlice, "Satya")
	mySlice = append(mySlice, "Sena")
	mySlice = append(mySlice, "John")

	log.Println(mySlice)

	var slice2 []int

	slice2 = append(slice2, 3)
	slice2 = append(slice2, 1)
	slice2 = append(slice2, 2)

	log.Println(slice2)

	sort.Ints(slice2)

	log.Println(slice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers[0:2])
}

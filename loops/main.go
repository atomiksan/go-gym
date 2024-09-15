package main

import "log"

func main() {
	// looping over a range
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// looping over an array
	animals := []string{"dog", "cat", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// using a blank identifier because we don't care about the index sometimes

	for _, animal := range animals {
		log.Println(animal)
	}

	// looping over a map

	pets := make(map[string]string)

	pets["dog"] = "fido"
	pets["cat"] = "puffy"
	pets["horse"] = "charlie"

	for petType, pet := range pets {
		log.Println(petType, pet)
	}

	// looping over a string

	line := "Once upon a blue  moon"
	for i, c := range line {
		log.Println(i, ":", c)
	}
}

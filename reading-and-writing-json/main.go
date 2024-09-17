package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
  [
    {
      "first_name": "Clark",
      "last_name": "Kent",
      "hair_color": "black",
      "has_dog": true
    },
    {
      "first_name": "Bruce",
      "last_name": "Wayne",
      "hair_color": "black",
      "has_dog": false
    }
  ]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var myHero []Person

	var h1 Person
	h1.FirstName = "Wally"
	h1.LastName = "West"
	h1.HairColor = "red"
	h1.HasDog = false

	myHero = append(myHero, h1)

	var h2 Person
	h2.FirstName = "Diana"
	h2.LastName = "Prince"
	h2.HairColor = "black"
	h2.HasDog = false

	myHero = append(myHero, h2)

	newJson, err := json.MarshalIndent(myHero, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}

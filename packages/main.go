package main

import (
	"log"

	"packages/helpers"
)

func main() {
	var myVar helpers.SomeType
	log.Println("Hello World!")

	myVar.TypeName = "Satya"
	log.Println(myVar.TypeName)
}

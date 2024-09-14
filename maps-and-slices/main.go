package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
}

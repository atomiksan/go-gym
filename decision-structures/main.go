package main

import "log"

func main() {
	var isTrue bool
	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is", cat)
	} else {
		log.Println("cat is not", cat)
	}

	myNum := 100
	isTrue = false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to false")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 100 and istrue is set to true")
	}

	animal := "tiger"

	switch animal {
	case "cat":
		log.Println("animal is", animal)
	case "dog":
		log.Println("animal is", animal)
	case "horse":
		log.Println("animal is", animal)
	default:
		log.Println("animal is set to something else")
	}
}

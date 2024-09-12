package main

import (
	"log"
)

var s = "seven"

func main() {
	s2 := "six"
	log.Println("s is ", s)
	log.Println("s2 is ", s2)

	s4, s5 := saySomething("xxx")
	log.Println("s4 is ", s4)
	log.Println("s5 is ", s5)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from saySomething func is ", s3)
	return s3, "world"
}

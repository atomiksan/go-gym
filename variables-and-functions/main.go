package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel World."
	i = 7
	fmt.Println(whatToSay)
	fmt.Println("i is set to ", i)

	whatWasSaid, theOtherthing := saySomething()

	fmt.Println("the fucntion saySomething said: ", whatWasSaid, theOtherthing)
}

func saySomething() (string, string) {
	return "Something", "else"
}

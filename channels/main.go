package main

import (
	"log"

	"github.com/atomiksan/go-projects/channels/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randNumber := helpers.RandNumber(numPool)
	intChan <- randNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

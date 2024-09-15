package helpers

import (
	"math/rand"
)

func RandNumber(n int) int {
	value := rand.Intn(n)
	return value
}

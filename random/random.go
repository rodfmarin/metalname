// Package random provides a simple interface for generating random numbers.
package random

import (
	"math/rand"
	"time"
)

// RandInt generate a new random int up to the max
func RandInt(max int) int {
	// Seed a random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

// Package random provides a simple interface for generating random numbers.
package random

import (
	"math/rand"
	"time"
)

// Number struct with ceil for generating rand nums
type Number struct {
	ceil int
}

// IRandomNumber interface for generating random numbers
type IRandomNumber interface {
	RandInt() int
}

// NewNumber generates a new random Number type
func NewNumber(ceil int) *Number {
	return &Number{ceil: ceil}
}

// RandInt generate a new random int
func (n Number) RandInt() int {
	// Seed a random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n.ceil)
}

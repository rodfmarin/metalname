package main

import (
	"fmt"

	"metalname/random"
)

func main() {
	// Create a random name generator
	nameGenerator, err := random.NewName()
	name, err := nameGenerator.Name()
	if err != nil {
		fmt.Printf("Failed to generate name: %v", err)
	}

	fmt.Printf("Staff Name: %s", name)
}

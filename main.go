package main

import (
	"fmt"

	"metalname/constants"
	"metalname/random"
)

func main() {
	// Create a name loader
	loader := constants.NewLoader()
	err := loader.LoadNames()
	if err != nil {
		panic(err)
	}
	// Create a randomizer
	r := random.NewNumber(len(loader.FirstNames))
	r2 := random.NewNumber(len(loader.LastNames))

	randFirst := r.RandInt()
	randLast := r2.RandInt()

	fmt.Printf("Staff Name: %s %s", loader.FirstNames[randFirst], loader.LastNames[randLast])

}

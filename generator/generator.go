// Package generator responsible for generating Names
package generator

import "metalname/constants"

// The generator type is responsible for generating names
type Generator struct {
	Loader constants.Loader
}

// New creates a new generator with the given name loader
func New(l constants.Loader) *Generator {
	return &Generator{
		Loader: l,
	}
}

// Generate generates a name by combining a random first name and last name
func (g *Generator) Generate() (string, error) {
	err := g.Loader.LoadNames()
	if err != nil {
		return "", err
	}
	firstName := g.Loader.FirstNames[0] // Replace with random selection logic
	lastName := g.Loader.LastNames[0]   // Replace with random selection logic
	return firstName + " " + lastName, nil
}

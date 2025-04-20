// Package random provides a way to generate items randomly.
package random

import (
	"fmt"

	"metalname/constants"
	"metalname/load"
	"metalname/schema"
)

// nameType is an enum to differentiate the different types of names
type nameType int

// Define the name types here (first name, last name, etc. - can be expanded to middle name, nickname, etc.)
const (
	first nameType = iota
	last
)

// String provides a string representation of the name type
func (n nameType) String() string {
	switch n {
	case first:
		return "First Name"
	case last:
		return "Last Name"
	default:
		return "Unknown Name Type"
	}
}

// nameMap adds a type for a map of name typess to a list of associated names
type nameMap map[nameType][]string

// Name contains 
type Name struct {
	storedNames nameMap	// storedNames is a map of the name type and 
}

// NewName creates a new Name loader instance
func NewName() (*Name, error) {
	n := &Name{
		storedNames: make(nameMap),
	}
	err := n.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to generate name generator: %v", err)
	}
	return n, nil
}

// names provides a way to safely retrieve the map of names
func (n *Name) names() nameMap {
	if n.storedNames == nil {
		n.storedNames = make(nameMap)
	}
	return n.storedNames
}

// Load reads files and stores the data
func (n *Name) Load() error {
	// Map of name type to the file that contains those names
	fileNames := map[nameType]string {
		first: constants.FirstNamesFile,
		last: constants.LastNamesFile,
	}
	// Load each name type
	for nameType, fileName := range fileNames {
		names, err := load.LoadFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to load %s: %v", nameType, err)
		}
		n.names()[nameType] = names
	}
	return nil
}

func (n *Name) randomName(nType nameType) (string, error) {
	names, ok := n.names()[nType]
	if !ok {
		err := n.Load()
		if err != nil {
			return "", fmt.Errorf("failed to retrieve a %s: %v", nType, err)
		}
	}
	numNames := len(names)
	if numNames == 0 {
		return "", fmt.Errorf("no names found for %s type", nType)
	}
	return names[RandInt(numNames)], nil
}

// FirstName retrieves a random first name from the list of names
func (n *Name) FirstName() (string, error) {
	return n.randomName(first)
}

// LastName retrieves a random last name from the list of names
func (n *Name) LastName() (string, error) {
	return n.randomName(last)
}

func (n *Name) Name() (*schema.Name, error) {
	firstName, err := n.FirstName()
	if err != nil {
		return nil, fmt.Errorf("failed to generate first name: %v", err)
	}
	lastName, err := n.LastName()
	if err != nil {
		return nil, fmt.Errorf("failed to generate last name: %v", err)
	}
	return &schema.Name{
		FirstName: firstName,
		LastName: lastName,
	}, nil
}
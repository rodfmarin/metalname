// Package constants responsible for constants like first name, last name, etc.
package constants

import (
	"bufio"
	"fmt"
	"os"
)

// The first names and last names files
const (
	FirstNamesFile = "./first_names.txt"
	LastNamesFile  = "./last_names.txt"
)

// ILoader an interface for loading names
type ILoader interface {
	LoadNames() error
}

// Loader is a struct that implements the ILoader interface
type Loader struct {
	FirstNames []string
	LastNames  []string
}

// NewLoader creates a new Loader instance
func NewLoader() *Loader {
	return &Loader{
		FirstNames: []string{},
		LastNames:  []string{},
	}
}

// LoadNames wraps loadFirstNames and loadLastNames to load all names from text files
func (l *Loader) LoadNames() error {
	fn, err := l.loadFirstNames()
	if err != nil {
		return fmt.Errorf("loading first names: %w", err)
	}
	ln, err := l.loadLastNames()
	if err != nil {
		return fmt.Errorf("loading last names: %w", err)
	}
	l.FirstNames = append(l.FirstNames, fn...)
	l.LastNames = append(l.LastNames, ln...)
	return nil
}

// loadFirstNames loads first names from a text file
func (l *Loader) loadFirstNames() ([]string, error) {
	names, err := load("first_names")
	if err != nil {
		return nil, fmt.Errorf("loading first names: %w", err)
	}
	return names, nil
}

// loadLastNames loads last names from a text file
func (l *Loader) loadLastNames() ([]string, error) {
	names, err := load("last_names")
	if err != nil {
		return nil, fmt.Errorf("loading last names: %w", err)
	}
	return names, nil
}

// load is the general purpose function that reads a file, given a file path
func load(nt string) ([]string, error) {

	var fpath string
	switch nt {
	case "first_names":
		{
			fpath = FirstNamesFile

		}
	case "last_names":
		{
			fpath = LastNamesFile
		}
	default:
		return nil, fmt.Errorf("unknown name type: %s", nt)
	}
	// Open the file, reading all lines
	file, err := os.Open(fpath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", fpath, err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fpath, err)
	}
	// Return the lines
	if len(lines) == 0 {
		return nil, fmt.Errorf("file %s is empty", fpath)
	}
	return lines, nil
}

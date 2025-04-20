// Package load is responsible for loading data from a source.
package load

import (
	"bufio"
	"fmt"
	"os"
)

// LoadFile loads the contents of a file into a slice of strings of each line of the file
func LoadFile(filePath string) ([]string, error) {
	// Open the file, reading all lines
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	// Verify the file is closed properly when done
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Failed to close file %s: %v", filePath, err)
		}
	}(file)

	// Read the file line by line and store the contents
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	// Return the lines
	if len(lines) == 0 {
		return nil, fmt.Errorf("file %s is empty", filePath)
	}
	return lines, nil
}
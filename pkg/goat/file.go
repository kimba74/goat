package goat

import (
	"fmt"
	"os"
	"strings"
)

func getFile(filename string) (*os.File, error) {
	// Check if provided filename is not empty
	if strings.TrimSpace(filename) == "" {
		return nil, fmt.Errorf("getFile: filename provided must not be empty")
	}

	// Check if filename is '-' for Stdin
	if strings.TrimSpace(filename) == "-" {
		// Ensure that Stdin is a named pipe
		if i, _ := os.Stdin.Stat(); i.Mode()&os.ModeNamedPipe == 0 {
			return nil, fmt.Errorf("getFile: stdin is not a named pipe")
		}
		return os.Stdin, nil
	}

	// Load the file for specified filename
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("getFile: %s", err)
	}
	return file, nil
}

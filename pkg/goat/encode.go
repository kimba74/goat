package goat

import (
	"fmt"
	"os"
	"strings"
)

// Encode values in the provided format
func Encode(format string, values ValueMap) error {
	switch strings.ToLower(format) {
	case "json":
		encodeJSON(os.Stdout, values)
	case "yaml":
		encodeYAML(os.Stdout, values)
	default:
		return fmt.Errorf("encoding values: output format '%s' is unknown, must be either json or yaml", format)
	}
	return nil
}

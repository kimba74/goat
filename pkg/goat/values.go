package goat

import (
	"fmt"
	"io/ioutil"
	"os"
)

func parseValues(file *os.File) (ValueMap, error) {
	// Read content of values file
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("parsing values: %s", err)
	}

	var val ValueMap

	// Try decoding the data as YAML
	val, err = decodeYAML(raw)
	if err == nil {
		return val, nil
	}

	// Try decoding the data as JSON
	val, err = decodeJSON(raw)
	if err == nil {
		return val, nil
	}

	// Error out if data was neither YAML nor JSON
	return nil, fmt.Errorf("parseValues: File '%s' is invalid or of unknown format. File must be either vaid yaml or json", file.Name())
}

// ParseValuesFile and return the content as a ValueMap
func ParseValuesFile(filename string) (ValueMap, error) {
	file, err := getFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ParseValuesFile: %s", err)
	}
	return parseValues(file)
}

// ParseValuesFiles and return their content as []ValueMap
func ParseValuesFiles(filename ...string) ([]ValueMap, error) {
	var vals []ValueMap
	// Parse if some filenames were provided
	if len(filename) > 0 {
		// Iterate over args provided and parse them
		for _, file := range filename {
			// Parse the filename
			val, err := ParseValuesFile(file)
			if err != nil {
				return nil, fmt.Errorf("ParseValueFiles: %s", err)
			}
			// Add the parsed ValueMap to the return values
			vals = append(vals, val)
		}
		// Return the slice of ValueMaps
		return vals, nil
	}
	//TODO Try processing Stdin if no arg was provided
	return nil, fmt.Errorf("ParseValueFiles: no argument specified")
}

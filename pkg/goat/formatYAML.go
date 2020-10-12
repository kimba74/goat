package goat

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

func decodeYAML(raw []byte) (ValueMap, error) {
	// Parse content of values file
	val := make(map[string]interface{})
	if err := yaml.Unmarshal(raw, &val); err != nil {
		return nil, fmt.Errorf("parsing values: %s", err)
	}
	return val, nil
}

func encodeYAML(w io.Writer, data ValueMap) {
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(3)
	encoder.Encode(data)
}

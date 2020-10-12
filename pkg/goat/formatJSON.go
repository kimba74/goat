package goat

import (
	"encoding/json"
	"fmt"
	"io"
)

func decodeJSON(raw []byte) (ValueMap, error) {
	// Parse content of values file
	val := make(map[string]interface{})
	if err := json.Unmarshal(raw, &val); err != nil {
		return nil, fmt.Errorf("parsing values: %s", err)
	}
	return val, nil
}

func encodeJSON(w io.Writer, data ValueMap) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "   ")
	encoder.Encode(data)
}

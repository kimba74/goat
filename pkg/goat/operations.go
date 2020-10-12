package goat

import (
	"fmt"
	"os"
	"text/template"

	"github.com/imdario/mergo"
)

// Apply value files to templates
func Apply(tmpl *template.Template, values ValueMap) error {
	// Check if template was provided
	if tmpl == nil {
		return fmt.Errorf("Apply: no template provided")
	}

	// Check if values were provided
	if len(values) <= 0 {
		return fmt.Errorf("Apply: no values provided")
	}

	// Execute the template with the provided values
	err := tmpl.Execute(os.Stdout, values)
	if err != nil {
		return fmt.Errorf("Apply %s", err)
	}

	return nil
}

// Merge ValueMaps and return as combined ValuesMap
func Merge(appendSlice bool, deepCopySlice bool, override bool, value ...ValueMap) (ValueMap, error) {
	var opt []func(*mergo.Config)
	if appendSlice {
		opt = append(opt, mergo.WithAppendSlice)
	}
	if deepCopySlice {
		opt = append(opt, mergo.WithSliceDeepCopy)
	}
	if override {
		opt = append(opt, mergo.WithOverride)
	}

	var val ValueMap = make(map[string]interface{})
	for _, v := range value {
		if err := mergo.Merge(&val, v, opt...); err != nil {
			return nil, err
		}
	}
	return val, nil
}

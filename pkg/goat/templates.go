package goat

import (
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig"
)

// ParseTemplateFile and return as Golang template
func ParseTemplateFile(filename string) (*template.Template, error) {
	// Get file for filename
	file, err := getFile(filename)
	if err != nil {
		return nil, fmt.Errorf("LoadTemplate: %s", err)
	}

	// Create new template for filename
	tmpl := template.New(filename)

	// Load 'sprig' functions
	tmpl.Funcs(sprig.TxtFuncMap())

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("LoadTemplate: %s", err)
	}

	t, err := tmpl.Parse(string(raw))
	if err != nil {
		return nil, fmt.Errorf("LoadTemplate: %s", err)
	}

	return t, nil
}

package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	tFile    string
	valFiles = []string{
		"my-val1.yaml",
		"my-val2.yaml",
		"my-val3.yaml",
	}
	yout = yaml.NewEncoder(os.Stdout)
)

// ValueMap representation of YAML file
type ValueMap map[string]interface{}

func init() {
	flag.StringVar(&tFile, "template", "", "path to go-template file")

	flag.Parse()
}

func main() {
	origin := []string{"a", "b", "c", "-", "a", "c", "d", "-", "a", "e", "g", "-", "f"}

	commap := make(map[string]bool)
	result := make([]string, 0)
	for _, v := range origin {
		i, ok := commap[v]
		if !ok {
			if v == "-" {
				commap[v] = i
			}
			result = append(result, v)
		}
	}

	fmt.Printf("origin: %v\n", origin)
	fmt.Printf("result: %v\n", result)

	fmt.Printf("slice: %v\n", origin[1:])
}

func tester(name string, values ...string) {

}

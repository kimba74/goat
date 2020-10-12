package goat

import "fmt"

// CheckMoreThanOnce checks if specified values exist more than once in input slice
func CheckMoreThanOnce(input []string, value ...string) error {
	// Create trigget map
	trigger := map[string]bool{}
	for _, v := range value {
		trigger[v] = true
	}
	// Check for duplicate trigger
	found := map[string]bool{}
	for _, in := range input {
		_, tOK := trigger[in]
		if tOK {
			_, fOK := found[in]
			if fOK {
				return fmt.Errorf("check_duplicate: found '%s' twice", in)
			}
			found[in] = true
		}
	}
	return nil
}

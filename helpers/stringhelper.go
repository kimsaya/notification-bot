package helpers

import "strings"

// StringContain compare string in list with input string
func StringContain(listString []string, key string) bool {
	for _, selected := range listString {
		if strings.Contains(selected, key) {
			return true
		}
	}
	return false
}

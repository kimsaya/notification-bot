package helpers

import (
	"log"
	"strconv"
	"strings"
)

// StringContains compare string in list with input string
func StringContains(listString []string, key string) bool {
	for _, selected := range listString {
		if strings.Contains(selected, key) {
			return true
		}
	}
	return false
}

func StringToInt64(input string) int64 {
	if i, err := strconv.ParseInt(input, 10, 64); err == nil {
		return i
	} else {
		log.Println(err)
		return -1
	}
}

func Int64ToString(input int64) string {
	return strconv.FormatInt(input, 10)
}

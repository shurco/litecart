package strutil

import "strings"

// ToSlice split string to array.
func ToSlice(s string, sep ...string) []string {
	if len(sep) > 0 {
		return strings.Split(s, sep[0])
	}
	return strings.Split(s, ",")
}

// ToAny split string to anu.
func ToAny(key ...string) []any {
	keys := make([]any, len(key))
	for i, v := range key {
		keys[i] = v
	}
	return keys
}

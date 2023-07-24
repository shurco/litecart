package strutil

import "strings"

// ToSlice split string to array.
func ToSlice(s string, sep ...string) []string {
	if len(sep) > 0 {
		return strings.Split(s, sep[0])
	}
	return strings.Split(s, ",")
}

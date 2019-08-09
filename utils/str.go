package utils

import "strings"

func RemoveSubString(input string, sub string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(sub, r) < 0 {
			return r
		}
		return -1
	}

	return strings.Map(filter, input)

}

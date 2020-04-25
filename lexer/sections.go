package lexer

import "strings"

var sections = []string{
	"Host",
	"Match",
}

func isSection(i string) bool {
	for _, val := range sections {
		if strings.ToLower(i) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

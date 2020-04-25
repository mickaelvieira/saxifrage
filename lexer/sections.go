package lexer

import "strings"

var sections = map[string]string{
	"host":  "Host",
	"match": "Match",
}

func isSection(i string) bool {
	k := strings.ToLower(i)
	_, ok := sections[k]
	return ok
}

func getSection(i string) string {
	k := strings.ToLower(i)
	v, _ := sections[k]
	return v
}

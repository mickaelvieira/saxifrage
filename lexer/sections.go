package lexer

var sections = []string{
	"Host",
	"Match",
}

func isSection(i string) bool {
	for _, val := range sections {
		if i == val {
			return true
		}
	}
	return false
}

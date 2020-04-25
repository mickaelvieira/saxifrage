package lexer

import "strings"

type section struct {
	ID   string
	Name string
}

var sections = []*section{
	{ID: "host", Name: "Host"},
	{ID: "match", Name: "Match"},
}

func isSection(i string) bool {
	for _, s := range sections {
		if strings.ToLower(i) == s.ID {
			return true
		}
	}
	return false
}

func getSection(i string) *section {
	for _, s := range sections {
		if strings.ToLower(i) == s.ID {
			return s
		}
	}
	return nil
}

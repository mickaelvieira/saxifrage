package lexer

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isHash(r rune) bool {
	return r == '#'
}

func isSeparator(r rune) bool {
	return r == '=' || r == ' '
}

func isDoubleQuote(r rune) bool {
	return r == '"'
}

func isEOF(r rune) bool {
	return r == eof
}

func isEOL(r rune) bool {
	return r == eol
}

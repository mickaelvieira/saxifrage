package lexer

// Type lexer token type
type Type int

// Lexer tokens
const (
	Illegal Type = iota
	Err
	EOF
	EOL
	Section
	Keyword
	Value
	Whitespace
	Comment
)

// Token is a token returned by the lexer
type Token struct {
	Type  Type
	Value string
}

// IsSection is the token a section?
func (t *Token) IsSection() bool {
	return t.Type == Section
}

// IsKeyword is the token a keyword?
func (t *Token) IsKeyword() bool {
	return t.Type == Keyword
}

// IsValue is the token a value?
func (t *Token) IsValue() bool {
	return t.Type == Value
}

// IsComment is the token a comment?
func (t *Token) IsComment() bool {
	return t.Type == Comment
}

// IsError is the token an error?
func (t *Token) IsError() bool {
	return t.Type == Err
}

// IsWhitespace is the token a whitespace?
func (t *Token) IsWhitespace() bool {
	return t.Type == Whitespace
}

// IsEOL is the token an EOL?
func (t *Token) IsEOL() bool {
	return t.Type == EOL
}

// IsEOF is the token an EOF?
func (t *Token) IsEOF() bool {
	return t.Type == EOF
}

// IsHostSection is the token a keyword with the value "Host"
func (t *Token) IsHostSection() bool {
	return t.IsSection() && t.Value == "Host"
}

// IsMatchSection is the token a keyword with the value "Match"
func (t *Token) IsMatchSection() bool {
	return t.IsSection() && t.Value == "Match"
}

// String returns the token as a string
func (t *Token) String() string {
	switch {
	case t.IsEOF():
		return "EOF"
	default:
		return t.Value
	}
}

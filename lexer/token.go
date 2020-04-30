package lexer

// TokenType lexer token type
type TokenType int

// Lexer tokens
const (
	Illegal    TokenType = iota // A illegal token is returned when an unexpected token is encounter
	Err                         // A token is returned when an error is found during lexing
	EOF                         // the end of the file
	EOL                         // the end of a line
	Section                     // either Host or Match section
	Keyword                     // a keyword (AddressFamily, BatchMode, BindAddress, etc...)
	Separator                   // either a space or equal sign to separate keywords and values
	Value                       // the value of a keyword or a section
	Whitespace                  // a sequence of blank characters, either spaces or tabs
	Comment                     // a comment
)

// Token is a token returned by the lexer.
// It has a type and a value representing
// the sequence of characters found in the input string
type Token struct {
	Type  TokenType
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

// IsSeparator is the token a separator?
func (t *Token) IsSeparator() bool {
	return t.Type == Separator
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

// IsIllegal is the token illegal?
func (t *Token) IsIllegal() bool {
	return t.Type == Illegal
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

// ToBytes returns the tokens as bytes
func (t *Token) ToBytes() []byte {
	return []byte(t.Value)
}

package keys

import (
	"errors"
	"strings"
)

// Type the type of the key, rsa, dsa, , etc...
type Type string

// Types list
const (
	RSA     Type = "rsa"
	DSA     Type = "dsa"
	ECDSA   Type = "ecdsa"
	ED25519 Type = "ed25519"
	INVALID Type = "invalid"
)

var types = []Type{RSA, DSA, ECDSA, ED25519}

// GetDefaultType returns the default type
func GetDefaultType() Type {
	return RSA
}

// TypesToString returns the list of key types
// separated by a comma
func TypesToString() string {
	s := make([]string, len(types))
	for i, t := range types {
		s[i] = string(t)
	}
	return strings.Join(s[:], ", ")
}

// Keys errors
var (
	ErrNotImplementedKeyType  = errors.New("This type of key is not yet implemented")
	ErrInvalidKeyType         = errors.New("Invalid key type. Type should be equal to rsa, dsa, ecdsa or ed25519")
	ErrPrivateKeyNotGenerated = errors.New("Private key must be generated before generating the public key")
)

// GetKeyType retrieves key type from user's input
func GetKeyType(i string) Type {
	for _, t := range types {
		if i == string(t) {
			return t
		}
	}
	return INVALID
}

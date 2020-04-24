package keys

import (
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
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
	ErrBitSizeNotSpecified    = errors.New("bitsize value was not set")
	ErrKeyOverrideNotAllowed  = errors.New("Overriding the key is not allowed")
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

// EncodeToPEM ...
// https://golang.org/pkg/encoding/pem/#Block
func EncodeToPEM(privateKey crypto.PrivateKey) ([]byte, error) {
	var der []byte
	var blk *pem.Block

	switch sk := privateKey.(type) {
	case *rsa.PrivateKey:
		der = x509.MarshalPKCS1PrivateKey(sk)
		blk = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: der,
		}
	case *ecdsa.PrivateKey:
		der, err := x509.MarshalECPrivateKey(sk)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: der,
		}
	case *dsa.PrivateKey:
		der, err := asn1.Marshal(sk.PublicKey)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: der,
		}
	case ed25519.PrivateKey:
		der, err := x509.MarshalPKCS8PrivateKey(sk)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{
			Type:  "OPENSSH PRIVATE KEY",
			Bytes: der,
		}
	default:
		return nil, fmt.Errorf("Invalid KEY type %v", sk)
	}

	return pem.EncodeToMemory(blk), nil
}

// type Options struct {
// 	Type       Type
// 	PrivateKey string
// 	PublicKey  string
// 	Directory  string
// }

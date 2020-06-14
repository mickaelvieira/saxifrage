package keys

import (
	"crypto"
	"crypto/dsa"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"sort"
	"strconv"
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

// Options keys options
type Options struct {
	KeyType    Type
	Host       string
	PassPhrase string
	PrivateKey string
	PublicKey  string
	Directory  string
	KeySize    interface{}
}

// Types contains the list of key types
var Types = []Type{RSA, DSA, ECDSA, ED25519}

// GetDefaultType returns the default type
func GetDefaultType() Type {
	return RSA
}

// TypesToString returns the list of key types
// separated by a comma
func TypesToString() string {
	s := make([]string, len(Types))
	for i, t := range Types {
		s[i] = string(t)
	}
	return strings.Join(s[:], ", ")
}

// Keys errors
var (
	ErrNotImplementedKeyType  = errors.New("This type of key is not yet implemented")
	ErrInvalidKeySize         = errors.New("Invalid key size")
	ErrInvalidKeyType         = errors.New("Invalid key type. Type should be equal to rsa, dsa, ecdsa or ed25519")
	ErrPrivateKeyNotGenerated = errors.New("Private key must be generated before generating the public key")
	ErrKeySizeNotValid        = errors.New("key size is not valid")
	ErrKeyOverrideNotAllowed  = errors.New("Overriding the key is not allowed")
)

// GetKeyType retrieves key type from user's input
func GetKeyType(i string) Type {
	for _, t := range Types {
		if i == string(t) {
			return t
		}
	}
	return INVALID
}

func sortKeySizeValues(v []string) []string {
	sort.Slice(v, func(i, j int) bool {
		i1, err := strconv.Atoi(v[i])
		if err != nil {
			panic(err)
		}
		i2, err := strconv.Atoi(v[j])
		if err != nil {
			panic(err)
		}
		return i1 > i2
	})
	return v
}

// EncodeToPEM encodes the key into PEM
func EncodeToPEM(privateKey crypto.PrivateKey, pwd string) ([]byte, error) {
	var der []byte
	var err error
	var blk *pem.Block

	switch sk := privateKey.(type) {
	case *rsa.PrivateKey:
		der = x509.MarshalPKCS1PrivateKey(sk)
		blk = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: der}
	case *ecdsa.PrivateKey:
		der, err = x509.MarshalECPrivateKey(sk)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{Type: "EC PRIVATE KEY", Bytes: der}
	case *dsa.PrivateKey:
		der, err = asn1.Marshal(sk.PublicKey)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{Type: "PUBLIC KEY", Bytes: der}
	case ed25519.PrivateKey:
		der, err = x509.MarshalPKCS8PrivateKey(sk)
		if err != nil {
			return nil, err
		}
		blk = &pem.Block{Type: "OPENSSH PRIVATE KEY", Bytes: der}
	default:
		return nil, fmt.Errorf("Invalid KEY type %v", sk)
	}

	if pwd != "" {
		blk, err = x509.EncryptPEMBlock(rand.Reader, blk.Type, blk.Bytes, []byte(pwd), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
	}

	return pem.EncodeToMemory(blk), nil
}

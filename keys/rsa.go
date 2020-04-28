package keys

import (
	"crypto/rand"
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

var rsaSizes = map[string]int{
	"64":   64,
	"128":  128,
	"256":  256,
	"512":  512,
	"1024": 1024,
	"2048": 2048,
}

// RSAKeySize RSA complexity implementation
type RSAKeySize struct{}

// GetDefault returns the default complexity
func (c *RSAKeySize) GetDefault() string {
	return "2048"
}

// GetValues returns the available complexity values
func (c *RSAKeySize) GetValues() []string {
	values := make([]string, len(rsaSizes))

	var i int
	for k := range rsaSizes {
		values[i] = k
		i++
	}

	return sortKeySizeValues(values)
}

// GetValue returns the complexity value
func (c *RSAKeySize) GetValue(s string) interface{} {
	v, ok := rsaSizes[s]
	if ok {
		return v
	}
	return nil
}

// RSAGenerator generates RSA private and public keys
type RSAGenerator struct {
	pwd string
	ks  interface{}
	pk  *rsa.PrivateKey
}

// GenPublicKey ...
func (g *RSAGenerator) GenPublicKey() ([]byte, error) {
	if g.pk == nil {
		return nil, ErrPrivateKeyNotGenerated
	}

	key, err := ssh.NewPublicKey(g.pk.Public())
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(key), nil
}

// GenPrivateKey ...
func (g *RSAGenerator) GenPrivateKey() ([]byte, error) {
	ks, ok := g.ks.(int)
	if !ok {
		return nil, ErrKeySizeNotValid
	}

	key, err := rsa.GenerateKey(rand.Reader, ks)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key, g.pwd)
}

package keys

import (
	"crypto/dsa"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

var dsaSizes = map[string]dsa.ParameterSizes{
	"1024": dsa.L1024N160,
	// "2048": dsa.L2048N224,
	"2048": dsa.L2048N256,
	"3072": dsa.L3072N256,
}

// DSAKeySize ECDSA complexity implementation
type DSAKeySize struct{}

// GetDefault returns the default complexity
func (c *DSAKeySize) GetDefault() string {
	return "2048"
}

// GetValues returns the available complexity values
func (c *DSAKeySize) GetValues() []string {
	values := make([]string, len(dsaSizes))

	var i int
	for k := range dsaSizes {
		values[i] = k
		i++
	}

	return sortKeySizeValues(values)
}

// GetValue returns the complexity value
func (c *DSAKeySize) GetValue(s string) interface{} {
	v, ok := dsaSizes[s]
	if ok {
		return v
	}
	return nil
}

// DSAGenerator generates RSA private and public keys
type DSAGenerator struct {
	pwd string
	ks  interface{}
	pk  *dsa.PrivateKey
}

// GenPublicKey ...
func (g *DSAGenerator) GenPublicKey() ([]byte, error) {
	if g.pk == nil {
		return nil, ErrPrivateKeyNotGenerated
	}

	key, err := ssh.NewPublicKey(&g.pk.PublicKey)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(key), nil
}

// GenPrivateKey ...
func (g *DSAGenerator) GenPrivateKey() ([]byte, error) {
	ks, ok := g.ks.(dsa.ParameterSizes)
	if !ok {
		return nil, ErrKeySizeNotValid
	}

	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err := dsa.GenerateParameters(params, rand.Reader, ks); err != nil {
		return nil, err
	}

	key := new(dsa.PrivateKey)
	key.PublicKey.Parameters = *params

	if err := dsa.GenerateKey(key, rand.Reader); err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key, g.pwd)
}

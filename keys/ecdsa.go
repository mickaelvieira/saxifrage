package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

var ecdsaSizes = map[string]elliptic.Curve{
	"256": elliptic.P256(),
	"384": elliptic.P384(),
	"521": elliptic.P521(),
}

// ECDSAKeySize ECDSA complexity implementation
type ECDSAKeySize struct{}

// GetDefault returns the default complexity
func (c *ECDSAKeySize) GetDefault() string {
	return "256"
}

// GetValues returns the available complexity values
func (c *ECDSAKeySize) GetValues() []string {
	values := make([]string, len(ecdsaSizes))

	var i int
	for k := range ecdsaSizes {
		values[i] = k
		i++
	}

	return sortKeySizeValues(values)
}

// GetValue returns the complexity value
func (c *ECDSAKeySize) GetValue(s string) interface{} {
	v, ok := ecdsaSizes[s]
	if ok {
		return v
	}
	return nil
}

// ECDSAGenerator generates ECDSA private and public keys
type ECDSAGenerator struct {
	pwd string
	ks  interface{}
	pk  *ecdsa.PrivateKey
}

// GenPublicKey ...
func (g *ECDSAGenerator) GenPublicKey() ([]byte, error) {
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
func (g *ECDSAGenerator) GenPrivateKey() ([]byte, error) {
	ks, ok := g.ks.(elliptic.Curve)
	if !ok {
		return nil, ErrKeySizeNotValid
	}

	key, err := ecdsa.GenerateKey(ks, rand.Reader)
	if err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key, g.pwd)
}

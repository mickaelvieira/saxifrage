package keys

import (
	"crypto/dsa"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

// DSAGenerator generates RSA private and public keys
type DSAGenerator struct {
	bitSize int
	pk      *dsa.PrivateKey
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
	params := new(dsa.Parameters)

	// see http://golang.org/pkg/crypto/dsa/#ParameterSizes
	if err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160); err != nil {
		return nil, err
	}

	key := new(dsa.PrivateKey)
	key.PublicKey.Parameters = *params

	if err := dsa.GenerateKey(key, rand.Reader); err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key)
}

package keys

import (
	"crypto/rand"
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
)

// RSAGenerator generates RSA private and public keys
type RSAGenerator struct {
	bitSize int
	pk      *rsa.PrivateKey
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
	if g.bitSize == 0 {
		return nil, ErrBitSizeNotSpecified
	}

	key, err := rsa.GenerateKey(rand.Reader, g.bitSize)
	if err != nil {
		return nil, err
	}

	err = key.Validate()
	if err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key)
}

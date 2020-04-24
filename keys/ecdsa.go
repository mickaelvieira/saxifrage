package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

// ECDSAGenerator generates RSA private and public keys
type ECDSAGenerator struct {
	bitSize int
	pk      *ecdsa.PrivateKey
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
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // @TODO We should allow other curses
	if err != nil {
		return nil, err
	}

	g.pk = key

	return EncodeToPEM(key)
}

package keys

import (
	"crypto/ed25519"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

// ED25519Generator generates RSA private and public keys
type ED25519Generator struct {
	bitSize int
	pk      ed25519.PublicKey
}

// GenPublicKey ...
func (g *ED25519Generator) GenPublicKey() ([]byte, error) {
	if g.pk == nil {
		return nil, ErrPrivateKeyNotGenerated
	}

	key, err := ssh.NewPublicKey(g.pk)
	if err != nil {
		return nil, err
	}

	return ssh.MarshalAuthorizedKey(key), nil
}

// GenPrivateKey ...
func (g *ED25519Generator) GenPrivateKey() ([]byte, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader) // @TODO We should allow other curses
	if err != nil {
		return nil, err
	}

	g.pk = publicKey

	return EncodeToPEM(privateKey)
}

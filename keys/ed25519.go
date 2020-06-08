package keys

import (
	"crypto/ed25519"
	"crypto/rand"

	"golang.org/x/crypto/ssh"
)

// ED25519KeySize ED25519 complexity implementation
type ED25519KeySize struct{}

// GetDefault returns the default complexity
func (c *ED25519KeySize) GetDefault() string {
	return ""
}

// GetValues returns the available complexity values
func (c *ED25519KeySize) GetValues() []string {
	return make([]string, 0)
}

// GetValue returns the complexity value
func (c *ED25519KeySize) GetValue(s string) interface{} {
	return nil
}

// ED25519Generator generates RSA private and public keys
type ED25519Generator struct {
	pwd string
	pk  ed25519.PublicKey
}

// GenPublicKey generates a ED25519 public key
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

// GenPrivateKey generates a ED25519 private key
func (g *ED25519Generator) GenPrivateKey() ([]byte, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	g.pk = publicKey

	return EncodeToPEM(privateKey, g.pwd)
}

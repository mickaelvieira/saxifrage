package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestED25519GenPrivateKey(t *testing.T) {
	gen := &ED25519Generator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)
}

func TestED25519GenPublicKeyWithoutPrivateKey(t *testing.T) {
	gen := &ED25519Generator{}
	_, err := gen.GenPublicKey()
	assert.Equal(t, ErrPrivateKeyNotGenerated, err)
}

func TestED25519GenPublicKey(t *testing.T) {
	gen := &ED25519Generator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)

	_, err = gen.GenPublicKey()
	assert.Nil(t, err)
}

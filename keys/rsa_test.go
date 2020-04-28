package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSAGenPrivateKeyNoBitSize(t *testing.T) {
	gen := &RSAGenerator{}
	_, err := gen.GenPrivateKey()
	assert.Equal(t, ErrKeySizeNotValid, err)
}

func TestRSAGenPrivateKey(t *testing.T) {
	gen := &RSAGenerator{ks: 64}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)
}

func TestRSAGenPublicKeyWithoutPrivateKey(t *testing.T) {
	gen := &RSAGenerator{}
	_, err := gen.GenPublicKey()
	assert.Equal(t, ErrPrivateKeyNotGenerated, err)
}

func TestRSAGenPublicKey(t *testing.T) {
	gen := &RSAGenerator{ks: 64}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)

	_, err = gen.GenPublicKey()
	assert.Nil(t, err)
}

package keys

import (
	"crypto/elliptic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECDSAGenPrivateKey(t *testing.T) {
	gen := &ECDSAGenerator{ks: elliptic.P256()}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)
}

func TestECDSAGenPublicKeyWithoutPrivateKey(t *testing.T) {
	gen := &ECDSAGenerator{}
	_, err := gen.GenPublicKey()
	assert.Equal(t, ErrPrivateKeyNotGenerated, err)
}

func TestECDSAGenPublicKey(t *testing.T) {
	gen := &ECDSAGenerator{ks: elliptic.P256()}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)

	_, err = gen.GenPublicKey()
	assert.Nil(t, err)
}

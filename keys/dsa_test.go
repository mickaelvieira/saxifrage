package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSAGenPrivateKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)
}

func TestDSAGenPublicKeyWithoutPrivateKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPublicKey()
	assert.Equal(t, ErrPrivateKeyNotGenerated, err)
}

func TestDSAGenPublicKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)

	_, err = gen.GenPublicKey()
	assert.Nil(t, err)
}

package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGenPrivateKeyNoBitSize(t *testing.T) {
// 	gen := &DSAGenerator{}
// 	_, err := gen.GenPrivateKey()
// 	assert.Equal(t, ErrBitSizeNotSpecified, err)
// }

func TestGenPrivateKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)
}

func TestGenPublicKeyWithoutPrivateKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPublicKey()
	assert.Equal(t, ErrPrivateKeyNotGenerated, err)
}

func TestGenPublicKey(t *testing.T) {
	gen := &DSAGenerator{}
	_, err := gen.GenPrivateKey()
	assert.Nil(t, err)

	_, err = gen.GenPublicKey()
	assert.Nil(t, err)
}

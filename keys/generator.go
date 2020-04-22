package keys

// Generator interface describing what is a generator
type Generator interface {
	GenPublicKey() ([]byte, error)
	GenPrivateKey() ([]byte, error)
	GetKeyType() Type
}

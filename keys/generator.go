package keys

// Generator interface describing what is a generator
type Generator interface {
	GenPublicKey() ([]byte, error)
	GenPrivateKey() ([]byte, error)
}

func GetGenerator(t Type) Generator {
	switch t {
	case RSA:
		return &RSAGenerator{bitSize: 4096}
	case DSA:
		return &DSAGenerator{}
	case ECDSA:
		return &ECDSAGenerator{}
	case ED25519:
		return &ED25519Generator{}
	default:
		panic("Invalid generator type")
	}
}

package keys

// Generator interface describing what is a generator
type Generator interface {
	GenPublicKey() ([]byte, error)
	GenPrivateKey() ([]byte, error)
}

// KeySize interface describing how to retrieve KeySize struct
type KeySize interface {
	GetDefault() string
	GetValues() []string
	GetValue(v string) interface{}
}

// GetKeySize returns the generator matching the type of key
func GetKeySize(t Type) KeySize {
	switch t {
	case RSA:
		return &RSAKeySize{}
	case ECDSA:
		return &ECDSAKeySize{}
	case ED25519:
		return &ED25519KeySize{}
	default:
		panic("Invalid generator type")
	}
}

// GetGenerator returns the generator matching the type of key
func GetGenerator(o *Options) Generator {
	switch o.KeyType {
	case RSA:
		return &RSAGenerator{ks: o.KeySize, pwd: o.PassPhrase}
	case ECDSA:
		return &ECDSAGenerator{ks: o.KeySize, pwd: o.PassPhrase}
	case ED25519:
		return &ED25519Generator{pwd: o.PassPhrase}
	default:
		panic("Invalid generator type")
	}
}

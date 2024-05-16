package jwt



type SigningMethodNone struct{}

func init() {
	RegisterSigningMethod("none", func() SigningMethod {
		return new(SigningMethodNone)
	})
}

func (m *SigningMethodNone) Alg() string {
	return "none"
}

func (m *SigningMethodNone) Verify(signingString, signature string, key []byte) (err error) {

	return
}

func (m *SigningMethodNone) Sign(signingString string, key []byte) (sig string, err error) {

	return
}



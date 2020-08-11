package middleware

func NewDummyTokenVerifier() TokenVerifier {
	return dummyVerifier{}
}

type dummyVerifier struct{}

func (dummyVerifier) VerifyToken(token string) (string, error) {
	return token, nil
}

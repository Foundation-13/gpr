package mdlwr

import "fmt"

const (
	FakeValidAuthToken = "123456789"
	FakeValidUserID = "123456789"
	FakeValidBearerAuthToken = "Bearer 123456789"
)

func NewFakeAuthTokenVerifier() AuthMdlwrTokenVerifier {
	return fakeAuthTokenVerifier{}
}

type fakeAuthTokenVerifier struct{}

func (fakeAuthTokenVerifier) VerifyToken(idToken string) (string, error) {
	if idToken == FakeValidAuthToken {
		return FakeValidUserID, nil
	} else {
		return "", fmt.Errorf("invalid token")
	}
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package middlewaremocks

import mock "github.com/stretchr/testify/mock"

// TokenVerifier is an autogenerated mock type for the TokenVerifier type
type TokenVerifier struct {
	mock.Mock
}

// VerifyToken provides a mock function with given fields: token
func (_m *TokenVerifier) VerifyToken(token string) (string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

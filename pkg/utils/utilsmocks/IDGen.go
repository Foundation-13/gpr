// Code generated by mockery v1.0.0. DO NOT EDIT.

package utilmocks

import mock "github.com/stretchr/testify/mock"

// IDGen is an autogenerated mock type for the IDGen type
type IDGen struct {
	mock.Mock
}

// StringID provides a mock function with given fields:
func (_m *IDGen) StringID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

package mock

import (
	"github.com/stretchr/testify/mock"
)

type CryptoMock struct {
	mock.Mock
}

func (m *CryptoMock) Encrypt(text string) (string, error) {
	args := m.Called(text)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).(string), resultError
}

func (m *CryptoMock) Decrypt(text string) (string, error) {
	args := m.Called(text)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).(string), resultError
}

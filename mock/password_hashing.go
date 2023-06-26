package mock

import (
	"github.com/stretchr/testify/mock"
)

type PassHashMock struct {
	mock.Mock
}

func (m *PassHashMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := m.Called(password, cost)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0).([]byte), resultError
}

func (m *PassHashMock) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := m.Called(hashedPassword, password)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

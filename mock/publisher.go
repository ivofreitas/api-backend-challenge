package mock

import (
	"github.com/stretchr/testify/mock"
)

type PublisherMock struct {
	mock.Mock
}

func (m *PublisherMock) Publish(str string) error {
	args := m.Called(str)

	var resultError error
	if args.Get(0) != nil {
		resultError = args.Get(0).(error)
	}

	return resultError
}

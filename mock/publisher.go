package mock

import (
	"github.com/stretchr/testify/mock"
)

type PublisherMock struct {
	mock.Mock
}

func (m *PublisherMock) Publish(str string) {
	m.Called(str)
}

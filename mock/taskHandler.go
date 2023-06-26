package mock

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type TaskHandlerMock struct {
	mock.Mock
}

func (m *TaskHandlerMock) Create(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

func (m *TaskHandlerMock) List(ctx context.Context, param interface{}) (interface{}, error) {
	args := m.Called(ctx, param)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	return args.Get(0), resultError
}

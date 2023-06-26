package mock

import "github.com/stretchr/testify/mock"

type JsonMock struct {
	mock.Mock
}

func (m *JsonMock) Marshal(v any) ([]byte, error) {
	args := m.Called(v)

	var resultError error
	if args.Get(1) != nil {
		resultError = args.Get(1).(error)
	}

	if result, ok := args.Get(0).([]byte); ok {
		return result, nil
	}

	return nil, resultError
}

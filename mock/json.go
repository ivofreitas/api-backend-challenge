package mock

import "github.com/stretchr/testify/mock"

type JsonMock struct {
	mock.Mock
}

func (m *JsonMock) Marshal(v any) ([]byte, error) {
	args := m.Called(v)

	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}

	return args.Get(0).([]byte), nil
}

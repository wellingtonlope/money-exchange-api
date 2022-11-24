package usecase

import (
	"github.com/stretchr/testify/mock"
)

type CreateMock struct {
	mock.Mock
}

func (m *CreateMock) Handle(input CreateInput) (Output, error) {
	args := m.Called(input)
	var result Output
	if args.Get(0) != nil {
		result = args.Get(0).(Output)
	}
	return result, args.Error(1)
}

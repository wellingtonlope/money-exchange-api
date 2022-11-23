package usecase

import (
	"github.com/stretchr/testify/mock"
)

type CreateMock struct {
	mock.Mock
}

func (m *CreateMock) Handle(input CreateInput) (CreateOutput, error) {
	args := m.Called(input)
	var result CreateOutput
	if args.Get(0) != nil {
		result = args.Get(0).(CreateOutput)
	}
	return result, args.Error(1)
}

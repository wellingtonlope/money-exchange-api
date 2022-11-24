package usecase

import "github.com/stretchr/testify/mock"

type GetByIDMock struct {
	mock.Mock
}

func (m *GetByIDMock) Handle(ID string) (Output, error) {
	args := m.Called(ID)
	var result Output
	if args.Get(0) != nil {
		result = args.Get(0).(Output)
	}
	return result, args.Error(1)
}

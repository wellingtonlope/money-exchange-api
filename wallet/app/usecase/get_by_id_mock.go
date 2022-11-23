package usecase

import "github.com/stretchr/testify/mock"

type GetByIDMock struct {
	mock.Mock
}

func (m *GetByIDMock) Handle(ID string) (GetByIDOutput, error) {
	args := m.Called(ID)
	var result GetByIDOutput
	if args.Get(0) != nil {
		result = args.Get(0).(GetByIDOutput)
	}
	return result, args.Error(1)
}

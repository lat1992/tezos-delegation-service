package mocks

import (
	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/stretchr/testify/mock"
)

type MockTzktService struct {
	mock.Mock
}

func (m *MockTzktService) GetDelegationsFrom(offset int) ([]structs.Delegation, error) {
	args := m.Called(offset)
	return args.Get(0).([]structs.Delegation), args.Error(1)
}

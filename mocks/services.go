package mocks

import (
	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/stretchr/testify/mock"
)

type MockTezosDelegationService struct {
	mock.Mock
}

func (m *MockTezosDelegationService) Start() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTezosDelegationService) Index(isInit bool) error {
	args := m.Called(isInit)
	return args.Error(0)
}

func (m *MockTezosDelegationService) GetDelegations(year string) ([]structs.Data, error) {
	args := m.Called(year)
	return args.Get(0).([]structs.Data), args.Error(1)
}

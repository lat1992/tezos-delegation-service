package mocks

import (
	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) Close() {
	m.Called()
}

func (m *MockDatabase) AddDelegations(delegations []structs.Delegation) error {
	args := m.Called(delegations)
	return args.Error(0)
}

func (m *MockDatabase) GetDelegationsCount() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockDatabase) GetDelegations() ([]structs.Data, error) {
	args := m.Called()
	return args.Get(0).([]structs.Data), args.Error(1)
}

func (m *MockDatabase) GetDelegationsByYear(year string) ([]structs.Data, error) {
	args := m.Called(year)
	return args.Get(0).([]structs.Data), args.Error(1)
}

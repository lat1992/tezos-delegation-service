package services

import (
	"testing"

	"github.com/lat1992/tezos-delegation-service/mocks"
	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/stretchr/testify/assert"
)

func TestNewTezosDelegation(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	mockTzkt := new(mocks.MockTzktService)

	td := NewTezosDelegation(mockDB, mockTzkt)

	assert.NotNil(t, td)
	assert.Equal(t, mockDB, td.database)
	assert.Equal(t, mockTzkt, td.tezos)
}

func TestTezosDelegation_Index(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	mockTzkt := new(mocks.MockTzktService)

	td := NewTezosDelegation(mockDB, mockTzkt)

	testCases := []struct {
		name          string
		isInit        bool
		dbCount       int
		delegations   []structs.Delegation
		databaseError error
		tzktError     error
	}{
		{
			name:    "success",
			isInit:  false,
			dbCount: 0,
			delegations: []structs.Delegation{
				{Timestamp: "2023-01-01", Sender: structs.Sender{Address: "tz1"}, Level: 1, Amount: 100},
			},
			databaseError: nil,
			tzktError:     nil,
		},
		{
			name:          "database error",
			isInit:        false,
			dbCount:       0,
			delegations:   []structs.Delegation{},
			databaseError: assert.AnError,
			tzktError:     nil,
		},
		{
			name:          "tzkt service error",
			isInit:        false,
			dbCount:       0,
			delegations:   []structs.Delegation{},
			databaseError: nil,
			tzktError:     assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDB.On("GetDelegationsCount").Return(tc.dbCount, tc.databaseError).Once()
			if tc.databaseError == nil {
				mockTzkt.On("GetDelegationsFrom", tc.dbCount).Return(tc.delegations, tc.tzktError).Once()
				if tc.tzktError == nil {
					mockDB.On("AddDelegations", tc.delegations).Return(tc.databaseError).Once()
				}
			}

			err := td.Index(tc.isInit)

			if tc.databaseError != nil || tc.tzktError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockDB.AssertExpectations(t)
			mockTzkt.AssertExpectations(t)
		})
	}
}

func TestTezosDelegation_GetDelegations(t *testing.T) {
	mockDB := new(mocks.MockDatabase)
	mockTzkt := new(mocks.MockTzktService)

	td := NewTezosDelegation(mockDB, mockTzkt)

	testCases := []struct {
		name          string
		year          string
		expectedData  []structs.Data
		expectedError error
	}{
		{
			name: "get all",
			year: "",
			expectedData: []structs.Data{
				{Timestamp: "2023-01-01", Amount: "100", Delegator: "tz1", Level: "1"},
			},
			expectedError: nil,
		},
		{
			name: "by year",
			year: "2023",
			expectedData: []structs.Data{
				{Timestamp: "2023-01-01", Amount: "100", Delegator: "tz1", Level: "1"},
			},
			expectedError: nil,
		},
		{
			name:          "database error",
			year:          "2023",
			expectedData:  nil,
			expectedError: assert.AnError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.year == "" {
				mockDB.On("GetDelegations").Return(tc.expectedData, tc.expectedError).Once()
			} else {
				mockDB.On("GetDelegationsByYear", tc.year).Return(tc.expectedData, tc.expectedError).Once()
			}

			data, err := td.GetDelegations(tc.year)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Nil(t, data)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedData, data)
			}

			mockDB.AssertExpectations(t)
		})
	}
}

package database

import (
	"testing"
	"time"

	"github.com/lat1992/tezos-delegation-service/structs"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
)

func setupMockDB(t *testing.T) (pgxmock.PgxPoolIface, *Store) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}

	store := &Store{
		postgres: mock,
	}

	return mock, store
}

func TestAddDelegations(t *testing.T) {
	mock, store := setupMockDB(t)
	defer mock.Close()

	delegations := []structs.Delegation{
		{
			Sender:    structs.Sender{Address: "addr1"},
			Amount:    100,
			Level:     1000,
			Timestamp: "2023-01-01T00:00:00Z",
		},
		{
			Sender:    structs.Sender{Address: "addr2"},
			Amount:    200,
			Level:     1001,
			Timestamp: "2023-01-02T00:00:00Z",
		},
	}

	mock.ExpectExec("INSERT INTO delegation").
		WithArgs("addr1", "100", "1000", pgxmock.AnyArg()).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectExec("INSERT INTO delegation").
		WithArgs("addr2", "200", "1001", pgxmock.AnyArg()).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := store.AddDelegations(delegations)
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetDelegationsCount(t *testing.T) {
	mock, store := setupMockDB(t)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"count"}).AddRow(2)
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(rows)

	count, err := store.GetDelegationsCount()
	assert.NoError(t, err)
	assert.Equal(t, 2, count)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetDelegations(t *testing.T) {
	mock, store := setupMockDB(t)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"delegator", "amount", "level", "timestamp"}).
		AddRow("addr1", "100", "1000", time.Now()).
		AddRow("addr2", "200", "1001", time.Now())

	mock.ExpectQuery("SELECT delegator, amount, block_high as level, timestamp FROM delegation").
		WillReturnRows(rows)

	delegations, err := store.GetDelegations()
	assert.NoError(t, err)
	assert.Len(t, delegations, 2)
	assert.Equal(t, "addr1", delegations[0].Delegator)
	assert.Equal(t, "100", delegations[0].Amount)
	assert.Equal(t, "1000", delegations[0].Level)
	assert.Equal(t, "addr2", delegations[1].Delegator)
	assert.Equal(t, "200", delegations[1].Amount)
	assert.Equal(t, "1001", delegations[1].Level)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetDelegationsByYear(t *testing.T) {
	mock, store := setupMockDB(t)
	defer mock.Close()

	rows := pgxmock.NewRows([]string{"delegator", "amount", "level", "timestamp"}).
		AddRow("addr1", "100", "1000", time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC))

	mock.ExpectQuery("SELECT delegator, amount, block_high as level, timestamp FROM delegation WHERE timestamp > \\$1 AND timestamp < \\$2").
		WithArgs("2022-01-01", "2022-12-31").
		WillReturnRows(rows)

	delegations, err := store.GetDelegationsByYear("2022")
	assert.NoError(t, err)
	assert.Len(t, delegations, 1)
	assert.Equal(t, "addr1", delegations[0].Delegator)
	assert.Equal(t, "100", delegations[0].Amount)
	assert.Equal(t, "1000", delegations[0].Level)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

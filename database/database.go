package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lat1992/tezos-delegation-service/structs"
)

type Database interface {
	Close()

	AddDelegations(delegations []structs.Delegation) error
	GetDelegationsCount() (int, error)
	GetDelegations() ([]structs.Data, error)
	GetDelegationsByYear(year string) ([]structs.Data, error)
}

type postgresInterface interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close()
}

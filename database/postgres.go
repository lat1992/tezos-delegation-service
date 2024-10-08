package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	postgres postgresInterface
}

func NewStore(host, port, database, user, password string) (*Store, error) {
	pg, err := pgxpool.New(context.Background(), fmt.Sprintf("postgres://%s:%s/%s?user=%s&password=%s", host, port, database, user, password))
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %w", err)
	}
	return &Store{
		postgres: pg,
	}, nil
}

func (s *Store) Close() {
	s.postgres.Close()
}

package storages

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	postgres *pgx.Conn
}

func NewStorage(host, port, database, user, password string) (*Storage, error) {
	pg, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s/%s?user=%s&password=%s", host, port, database, user, password))
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %w", err)
	}
	return &Storage{
		postgres: pg,
	}, nil
}

func (s *Storage) Close() {
	if err := s.postgres.Close(context.Background()); err != nil {
		slog.Error("error when closing postgres connection", "error", err)
	}
}

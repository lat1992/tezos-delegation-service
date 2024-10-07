package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/lat1992/tezos-delegation-service/structs"
)

// AddDelegations inserts multiple delegation records into the database
// It uses a transaction and batch insert for efficiency
func (s *Store) AddDelegations(delegations []structs.Delegation) error {
	ctx := context.Background()

	tx, err := s.postgres.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}
	for _, d := range delegations {
		timestamp, err := time.Parse(time.RFC3339, d.Timestamp)
		if err != nil {
			return fmt.Errorf("parse timestamp %s: %w", d.Timestamp, err)
		}
		batch.Queue(
			"INSERT INTO delegation (delegator, amount, block_high, timestamp) VALUES ($1, $2, $3, $4)",
			d.Sender.Address,
			d.Amount,
			strconv.Itoa(d.Level),
			timestamp,
		)
	}

	br := tx.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < batch.Len(); i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("execute batch insert: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

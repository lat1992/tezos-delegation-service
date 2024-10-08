package database

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/lat1992/tezos-delegation-service/structs"
)

// AddDelegations inserts multiple delegation records into the database
func (s *Store) AddDelegations(delegations []structs.Delegation) error {
	ctx := context.Background()

	for _, d := range delegations {
		timestamp, err := time.Parse(time.RFC3339, d.Timestamp)
		if err != nil {
			return fmt.Errorf("parse timestamp %s: %w", d.Timestamp, err)
		}

		if _, err := s.postgres.Exec(
			ctx,
			"INSERT INTO delegation (delegator, amount, block_high, timestamp) VALUES ($1, $2, $3, $4)",
			d.Sender.Address,
			strconv.Itoa(d.Amount),
			strconv.Itoa(d.Level),
			timestamp,
		); err != nil {
			return fmt.Errorf("postgres sql exec error: %w", err)
		}
	}

	return nil
}

func (s *Store) GetDelegationsCount() (int, error) {
	ctx := context.Background()

	var count int
	if err := s.postgres.QueryRow(ctx, "SELECT COUNT(*) FROM delegation").Scan(&count); err != nil {
		return 0, fmt.Errorf("count delegations: %w", err)
	}

	return count, nil
}

func (s *Store) GetDelegations() ([]structs.Data, error) {
	rows, err := s.postgres.Query(context.Background(), "SELECT delegator, amount, block_high as level, timestamp FROM delegation ORDER BY timestamp DESC")
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	var delegations []structs.Data
	for rows.Next() {
		var delegation structs.Data
		var timestamp time.Time
		if err = rows.Scan(&delegation.Delegator, &delegation.Amount, &delegation.Level, &timestamp); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		delegation.Timestamp = timestamp.Format(time.RFC3339)
		delegations = append(delegations, delegation)
	}

	return delegations, nil
}

func (s *Store) GetDelegationsByYear(year string) ([]structs.Data, error) {
	begin := year + "-01-01"
	end := year + "-12-31"

	rows, err := s.postgres.Query(context.Background(), "SELECT delegator, amount, block_high as level, timestamp FROM delegation WHERE timestamp > $1 AND timestamp < $2 ORDER BY timestamp DESC", begin, end)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	var delegations []structs.Data
	for rows.Next() {
		var delegation structs.Data
		var timestamp time.Time
		if err = rows.Scan(&delegation.Delegator, &delegation.Amount, &delegation.Level, &timestamp); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		delegation.Timestamp = timestamp.Format(time.RFC3339)
		delegations = append(delegations, delegation)
	}

	return delegations, nil
}

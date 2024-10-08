package services

import (
	"fmt"
	"time"

	"github.com/lat1992/tezos-delegation-service/database"
	"github.com/lat1992/tezos-delegation-service/external"
	"github.com/lat1992/tezos-delegation-service/structs"
)

type TezosDelegation struct {
	tezos    external.TzktService
	database database.Database
}

func NewTezosDelegation(db database.Database, tezosClient external.TzktService) *TezosDelegation {
	return &TezosDelegation{
		tezos:    tezosClient,
		database: db,
	}
}

func (td *TezosDelegation) Start() error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := td.Index(false); err != nil {
				return fmt.Errorf("indexing error: %w", err)
			}
		}
	}
}

func (td *TezosDelegation) Index(isInit bool) error {
	for {
		offset, err := td.database.GetDelegationsCount()
		if err != nil {
			return fmt.Errorf("get delegation count: %w", err)
		}
		delegations, err := td.tezos.GetDelegationsFrom(offset)
		if err != nil {
			return fmt.Errorf("get delegations from offset %d: %w", offset, err)
		}
		if len(delegations) > 0 {
			if err = td.database.AddDelegations(delegations); err != nil {
				return fmt.Errorf("add delegation to database: %w", err)
			}
		} else {
			return nil
		}
		if !isInit {
			return nil
		}
	}
}

func (td *TezosDelegation) GetDelegations(year string) ([]structs.Data, error) {
	if year == "" {
		data, err := td.database.GetDelegations()
		if err != nil {
			return nil, fmt.Errorf("get delegation by year error :%w", err)
		}
		return data, nil
	}
	data, err := td.database.GetDelegationsByYear(year)
	if err != nil {
		return nil, fmt.Errorf("get delegation by year error :%w", err)
	}
	return data, nil
}

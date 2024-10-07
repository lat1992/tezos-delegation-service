package services

import (
	"github.com/lat1992/tezos-delegation-service/database"
	"github.com/lat1992/tezos-delegation-service/external"
	"github.com/lat1992/tezos-delegation-service/structs"
)

type TezosDelegation struct {
	tezos    external.TzktService
	database database.Database
}

func NewTezosDelegation(tezosClient external.TzktService) *TezosDelegation {
	return &TezosDelegation{
		tezos: tezosClient,
	}
}

func (td *TezosDelegation) Init() {
	delegations := td.tezos.GetAllDelegations()
	if len(delegations) > 0 {
		td.database
	}
}

func (td *TezosDelegation) GetDelegations() ([]structs.Data, error) {
	return nil, nil
}

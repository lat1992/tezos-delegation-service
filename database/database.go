package database

import "github.com/lat1992/tezos-delegation-service/structs"

type Database interface {
	Close()

	AddDelegations(delegations []structs.Delegation)
}

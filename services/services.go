package services

import "github.com/lat1992/tezos-delegation-service/structs"

type TezosDelegationService interface {
	GetDelegations() ([]structs.Data, error)
}

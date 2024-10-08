package services

import "github.com/lat1992/tezos-delegation-service/structs"

type TezosDelegationService interface {
	Start() error
	Index(isInit bool) error
	GetDelegations(year string) ([]structs.Data, error)
}

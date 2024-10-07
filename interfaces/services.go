package interfaces

import "github.com/lat1992/tezos-delegation-service/structures"

type TezosDelegationService interface {
	GetDatas() ([]structures.Data, error)
}

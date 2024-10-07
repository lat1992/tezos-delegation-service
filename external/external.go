package external

import "github.com/lat1992/tezos-delegation-service/structs"

type TzktService interface {
	GetAllDelegations() []structs.Delegation

	GetDelegationsFrom(offset int) ([]structs.Delegation, error)
}

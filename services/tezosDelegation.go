package services

import (
	"github.com/lat1992/tezos-delegation-service/structures"
)

type TezosDelegation struct {
}

func NewTezosDelegation() *TezosDelegation {
	return &TezosDelegation{}
}

func (*TezosDelegation) GetDatas() ([]structures.Data, error) {
	return nil, nil
}

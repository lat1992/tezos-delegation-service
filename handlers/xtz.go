package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tezos-delegation-service/services"
)

func delegations(tds services.TezosDelegationService) func(c *gin.Context) {
	return func(c *gin.Context) {
		datas, err := tds.GetDelegations()
		if err != nil {
			c.Error(fmt.Errorf("error when getting data: %w", err))
			return
		}
		c.JSON(200, gin.H{"data": datas})
	}
}

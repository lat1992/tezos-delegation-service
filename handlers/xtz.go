package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tezos-delegation-service/services"
)

func delegations(tds services.TezosDelegationService) func(c *gin.Context) {
	return func(c *gin.Context) {
		datas, err := tds.GetDelegations("")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": datas})
	}
}

func delegationsByYear(tds services.TezosDelegationService) func(c *gin.Context) {
	return func(c *gin.Context) {
		year := c.Param("year")
		if year != "" {
			if _, err := strconv.Atoi(year); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "parameter invalid"})
				return
			}
		}
		datas, err := tds.GetDelegations(year)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": datas})
	}
}

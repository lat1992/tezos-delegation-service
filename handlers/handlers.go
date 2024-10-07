package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lat1992/tezos-delegation-service/services"
)

func GetRouter(tds services.TezosDelegationService) *gin.Engine {
	router := gin.Default()

	router.GET("/", health)
	router.GET("/health", health)

	xtz := router.Group("/xtz")
	xtz.POST("/delegations/:year", delegations(tds))

	return router
}

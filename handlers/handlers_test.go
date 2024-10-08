package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lat1992/tezos-delegation-service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetRouter(t *testing.T) {
	mockTds := new(mocks.MockTezosDelegationService)
	router := GetRouter(mockTds)

	assert.NotNil(t, router)
	assert.IsType(t, &gin.Engine{}, router)

	routes := router.Routes()
	assert.Len(t, routes, 4)

	routePaths := []string{"/", "/health", "/xtz/delegations", "/xtz/delegations/:year"}
	for _, route := range routes {
		assert.Contains(t, routePaths, route.Path)
	}
}

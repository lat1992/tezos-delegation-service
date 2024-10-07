package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.Status(http.StatusOK)
}

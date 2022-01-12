package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

func NewHealthCtrl() (*healthCtrl, error) {
	return &healthCtrl{}, nil
}

func (h *healthCtrl) HealthEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}

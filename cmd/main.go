package main

import (
	"github.com/gin-gonic/gin"
	httpv1 "github.com/quyenphamkhac/saga-ecommerce-ordersvc/transport/http/v1"
)

func main() {
	r := gin.Default()
	healthCtrl, err := httpv1.NewHealthCtrl()
	if err != nil {
		panic(err)
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthCtrl.HealthEndpoint)
	}
	r.Run()
}

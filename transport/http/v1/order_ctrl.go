package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/usecase"
)

type orderCtrl struct {
	usecase *usecase.OrderUsecase
}

func NewOrderCtrl(usecase *usecase.OrderUsecase) (*orderCtrl, error) {
	return &orderCtrl{
		usecase: usecase,
	}, nil
}

func (o *orderCtrl) PlaceOrderEndpoint(c *gin.Context) {
}

func (o *orderCtrl) CancelOrderEndpoint(c *gin.Context) {
}

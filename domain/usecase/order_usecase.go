package usecase

import (
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/dto"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
)

type OrderUsecase interface {
	PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error)
	CancelOrder(data *dto.CancelOrderDto) (*model.Order, error)
}

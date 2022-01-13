package usecase

import (
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/dto"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/repository"
)

type OrderUsecase interface {
	PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error)
	CancelOrder(data *dto.CancelOrderDto) (*model.Order, error)
}

type orderUsecaseImpl struct {
	repo repository.OrderRepository
}

func NewOrderUsecaseImpl(repo repository.OrderRepository) *orderUsecaseImpl {
	return &orderUsecaseImpl{
		repo: repo,
	}
}

func (u *orderUsecaseImpl) PlaceOrder(data *dto.PlaceOrderDto) (*model.Order, error) {
	return nil, nil
}

func (u *orderUsecaseImpl) CancelOrder(data *dto.CancelOrderDto) (*model.Order, error) {
	return nil, nil
}

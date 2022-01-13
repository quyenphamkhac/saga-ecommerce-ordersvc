package repository

import (
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/dto"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
)

type BaseRepository interface {
	Find(search string) ([]model.Order, error)
	FindById(id string) (*model.Order, error)
	FindByIds(ids []string) ([]model.Order, error)
	Insert(data *dto.InsertOrderDto) (*model.Order, error)
	Update(id string, data *dto.UpdateOrderDto) (*model.Order, error)
	Delete(id string) (*model.Order, error)
}

type OrderRepository interface {
	BaseRepository
}

type OrderItemRepository interface {
	BaseRepository
}

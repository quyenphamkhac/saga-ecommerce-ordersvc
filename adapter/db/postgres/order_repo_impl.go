package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/dto"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
)

type orderRepoImpl struct {
	pg *pg.DB
}

func NewOrderRepoImpl(pgDB *pg.DB) *orderRepoImpl {
	return &orderRepoImpl{
		pg: pgDB,
	}
}

func (r *orderRepoImpl) Find(search string) ([]model.Order, error) {
	var orders []model.Order
	err := r.pg.Model(&orders).Select()
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (r *orderRepoImpl) FindById(id string) (*model.Order, error) {
	order := &model.Order{}
	r.pg.Model(order).WherePK()
	return nil, nil
}

func (r *orderRepoImpl) FindByIds(ids []string) ([]model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) Insert(data *dto.InsertOrderDto) (*model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) Update(id string, data *dto.UpdateOrderDto) (*model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) Delete(id string) (*model.Order, error) {
	return nil, nil
}

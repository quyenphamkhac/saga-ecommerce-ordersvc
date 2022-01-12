package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/dto"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
)

type orderRepoImpl struct {
	conn *pg.Conn
}

func NewOrderRepoImpl(postgresConn *pg.Conn) (*orderRepoImpl, error) {
	return &orderRepoImpl{
		conn: postgresConn,
	}, nil
}

func (r *orderRepoImpl) Find(search string) ([]model.Order, error) {
	return nil, nil
}

func (r *orderRepoImpl) FindById(id string) (*model.Order, error) {
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

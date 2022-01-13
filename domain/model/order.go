package model

import "time"

type Order struct {
	Id         int64
	CustomerId int64
	OrderItems []*OrderItem `pg:"rel:has-many"`
	Total      int16
	Status     string
	CreatedAt  time.Time `pg:"default:now()"`
	UpdatedAt  time.Time `pg:"default:now()"`
}

type OrderItem struct {
	ProductId   int64 `pg:",pk"`
	OrderId     int64 `pg:",pk"`
	ProductName string
	Quantity    int16
	Price       int16
}

package model

import "time"

type Order struct {
	Id         int64
	CustomerId int64
	OrderItems []*OrderItem `pg:"rel:has-many"`
	Total      int16
	CreatedAt  time.Time `pg:"default:now()"`
	UpdatedAt  time.Time `pg:"default:now()"`
}

type OrderItem struct {
	Id          int64
	ProductId   int64
	ProductName string
	Quantity    int16
	Price       int16
}

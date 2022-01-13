package dto

type InsertOrderDto struct{}

type UpdateOrderDto struct{}

type PlaceOrderDto struct {
	CustomerId int                   `json:"customer_id" binding:"required"`
	Total      int                   `json:"total" binding:"required"`
	OrderItems []*InsertOrderItemDto `json:"order_items" binding:"required,dive"`
}

type InsertOrderItemDto struct {
	ProductId   int    `json:"product_id" binding:"required"`
	ProductName string `json:"product_name" binding:"max=20,required"`
	Quantity    int    `json:"quantity" binding:"max=50,required"`
	Price       int    `json:"price" binding:"required"`
}

type CancelOrderDto struct{}

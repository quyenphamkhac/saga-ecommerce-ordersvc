package transport

import "github.com/quyenphamkhac/saga-ecommerce-ordersvc/config"

type mqTransport struct {
	cfg *config.Config
}

func NewMQTransport(cfg *config.Config) *mqTransport {
	return &mqTransport{
		cfg: cfg,
	}
}

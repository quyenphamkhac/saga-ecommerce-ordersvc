package app

type orderSvcApp struct {
}

type BaseService interface {
	Init() error
	Run() error
}

func NewOrderSvcApp() (*orderSvcApp, error) {
	return nil, nil
}

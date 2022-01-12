package app

type orderSvcApp struct {
}

type Service interface {
	Init() error
	Run() error
}

func NewOrderSvcApp() (*orderSvcApp, error) {
	return nil, nil
}

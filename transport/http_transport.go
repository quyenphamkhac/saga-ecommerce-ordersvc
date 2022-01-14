package transport

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/adapter/datastore/postgres"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/pkg/postgresql"

	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/config"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/usecase"
	httpv1 "github.com/quyenphamkhac/saga-ecommerce-ordersvc/endpoint/http/v1"
	httpmdw "github.com/quyenphamkhac/saga-ecommerce-ordersvc/middleware/http"
)

type httpServer struct {
	cfg *config.Config
}

func NewHttpServer(cfg *config.Config) *httpServer {
	return &httpServer{
		cfg: cfg,
	}
}

func (s *httpServer) Run(addr string) {
	db := postgresql.NewPostgresqlDBConn(s.cfg)
	defer db.Close()
	err := postgresql.CreatePostgresqlDBSchema(db)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(httpmdw.ErrorsMiddleware(gin.ErrorTypeAny))
	orderRepository := postgres.NewOrderRepoImpl(db)
	orderUsecase := usecase.NewOrderUsecaseImpl(orderRepository)

	healthCtrl := httpv1.NewHealthCtrl()
	orderCtrl := httpv1.NewOrderCtrl(orderUsecase)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthCtrl.HealthEndpoint)
		v1.POST("/orders", orderCtrl.PlaceOrderEndpoint)
	}
	go func() {
		if err := r.Run(addr); err != nil {
			log.Println("run http server failed")
		}
	}()
}

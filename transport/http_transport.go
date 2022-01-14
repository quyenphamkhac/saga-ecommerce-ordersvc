package transport

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/adapter/db/postgres"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/config"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/usecase"
	httpv1 "github.com/quyenphamkhac/saga-ecommerce-ordersvc/endpoint/http/v1"
	httpmdw "github.com/quyenphamkhac/saga-ecommerce-ordersvc/middleware/http"
)

var initDBConnOnce sync.Once

type httpServer struct {
	cfg *config.Config
}

func NewHttpServer(cfg *config.Config) *httpServer {
	return &httpServer{
		cfg: cfg,
	}
}

func (s *httpServer) Run(addr string) {
	db := initDBConn(s.cfg)
	defer db.Close()
	err := createDBSchema(db)
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

func initDBConn(cfg *config.Config) *pg.DB {
	var conn *pg.DB
	initDBConnOnce.Do(func() {
		conn = pg.Connect(&pg.Options{
			User:     cfg.PostgresDB.User,
			Password: cfg.PostgresDB.Password,
			Addr:     cfg.PostgresDB.Addr,
			Database: cfg.PostgresDB.Database,
		})
	})
	return conn
}

func createDBSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Order)(nil),
		(*model.OrderItem)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

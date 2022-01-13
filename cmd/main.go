package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/quyenphamkhac/saga-ecommerce-ordersvc/domain/model"
	httpv1 "github.com/quyenphamkhac/saga-ecommerce-ordersvc/transport/http/v1"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Addr:     "localhost:5433",
		Database: "ordersvc",
	})
	defer db.Close()

	err := migrateDb(db)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	healthCtrl, err := httpv1.NewHealthCtrl()
	if err != nil {
		panic(err)
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/health", healthCtrl.HealthEndpoint)
	}
	r.Run()
}

func migrateDb(db *pg.DB) error {
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

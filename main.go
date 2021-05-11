package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/saaitt/go-crud/handler"
	"github.com/saaitt/go-crud/model"
	"github.com/saaitt/go-crud/service"
	"github.com/saaitt/go-crud/sql"
	"os"
)

func main() {
	pass := os.Getenv("DB_PASSWORD")
	url := fmt.Sprintf("host=localhost port=5432 user=go dbname=go_crud_db password=%s sslmode=disable", pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Product{})
	defer db.Close()

	e := echo.New()
	product := handler.ProductHandler{
		Service: service.ProductService{
			Repo: &sql.ProductRepo{
				DB: db,
			},
		},
	}
	e.GET("/:page_no", product.List)
	e.POST("/", product.Create)
	e.Logger.Fatal(e.Start(":9876"))
}

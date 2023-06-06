package main

import (
	"lpfigueiredo/go-crud/product"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}

func main() {
	godotenv.Load()

	db := initDB()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindByID)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

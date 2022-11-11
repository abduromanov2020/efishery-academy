package config

import (
	"ecommerce-project/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

}

func AutoMigrate() {
	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Products{})
	DB.AutoMigrate(&entity.Cart{})
	DB.AutoMigrate(&entity.Cart_Product{})
	DB.AutoMigrate(&entity.Payment{})
	DB.AutoMigrate(&entity.Category{})
	DB.AutoMigrate(&entity.Product_Detail{})
}

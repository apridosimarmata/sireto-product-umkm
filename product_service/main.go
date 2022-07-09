package main

import (
	"fmt"
	Config "product_services/Config"
	Model "product_services/Model"
	Routes "product_services/Routes"

	"github.com/jinzhu/gorm"
)

func main() {
	var err error
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(
		&Model.Product{},
		&Model.ProductCategory{},
		&Model.ProductSubCategory{},
	)

	r := Routes.SetupRouter()

	r.Run(":8082")
}

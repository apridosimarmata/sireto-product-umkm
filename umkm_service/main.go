package main

import (
	"fmt"
	"umkm_service/Config"
	"umkm_service/Models"
	"umkm_service/Routes.go"

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
		&Models.UMKM{},
		&Models.UMKMCategory{},
		&Models.UMKMSubCategory{},
	)
	Config.DB.Model(&Models.UMKMSubCategory{}).AddForeignKey("category_id", "umkm_categories(id)", "RESTRICT", "RESTRICT")
	r := Routes.SetupRouter()

	r.Run()
}

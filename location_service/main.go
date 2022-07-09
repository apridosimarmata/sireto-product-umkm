package main

import (
	"fmt"
	"location_service/Config"
	"location_service/Models"
	"location_service/Routes.go"

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
		&Models.Province{},
		&Models.City{},
	)

	Config.DB.Model(&Models.City{}).AddForeignKey("province_id", "provinces(id)", "RESTRICT", "RESTRICT")

	r := Routes.SetupRouter()

	r.Run(":8081")
}

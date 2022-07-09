package Models

import (
	_ "github.com/go-sql-driver/mysql"

	"location_service/Config"
)

func CreateProvince(province *Province) (err error) {
	if err = Config.DB.Create(province).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProvinces(provinces *[]Province) (err error) {
	if err = Config.DB.Find(&provinces).Error; err != nil {
		return err
	}
	return nil
}

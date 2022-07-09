package Models

import "location_service/Config"

func CreateCity(city *City) (err error) {
	if err = Config.DB.Create(city).Error; err != nil {
		return err
	}
	return nil
}

func GetCitiesByProvinceId(cities *[]City, province_id string) (err error) {
	if err = Config.DB.Where("province_id = ?", province_id).Find(cities).Error; err != nil {
		return err
	}

	return nil
}

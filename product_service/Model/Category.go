package model

import (
	Config "product_services/Config"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCategory(category *ProductCategory) (err error) {
	if err = Config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func CreateSubCategory(subCategory *ProductSubCategory) (err error) {
	if err = Config.DB.Create(subCategory).Error; err != nil {
		return err
	}
	return nil
}

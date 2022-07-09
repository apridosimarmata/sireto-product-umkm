package model

import (
	_ "github.com/go-sql-driver/mysql"

	Config "product_services/Config"
)

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductById(product *Product, productId string) (err error) {
	if err = Config.DB.Where("id = ?", productId).Find(product).Error; err != nil {
		return err
	}

	return nil
}

func GetProductByName(products *[]Product, productName string, categoryId string) (err error) {
	if err = Config.DB.Where("name like ? and category_id = ?", "%"+productName+"%", categoryId).Find(products).Error; err != nil {
		return err
	}

	return nil
}

func GetProductByMerchantId(products *[]Product, merchantId string) (err error) {
	if err = Config.DB.Where("merchant_id = ?", merchantId).Find(products).Error; err != nil {
		return err
	}
	return nil
}

func GetBestSellerProduct(products *[]Product) (err error) {
	if err = Config.DB.Order("sold asc").Limit(7).Find(products).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product) (err error) {
	Config.DB.Save(product)
	return nil
}

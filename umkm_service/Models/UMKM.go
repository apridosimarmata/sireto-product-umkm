package Models

import (
	_ "github.com/go-sql-driver/mysql"

	"umkm_service/Config"
)

func CreateUMKM(umkm *UMKM) (err error) {
	if err = Config.DB.Create(umkm).Error; err != nil {
		return err
	}
	return nil
}

func CreateUMKMCategory(umkmCategory *UMKMCategory) (err error) {
	if err = Config.DB.Create(umkmCategory).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUMKMCategories(umkmCategories *[]UMKMCategory) (err error) {
	if err = Config.DB.Find(umkmCategories).Error; err != nil {
		return err
	}
	return nil
}

func CreateUMKMSubCategory(umkmSubCategory *UMKMSubCategory) (err error) {
	if err = Config.DB.Create(umkmSubCategory).Error; err != nil {
		return err
	}
	return nil
}

func GetUMKMSubCategoriesByCategoryId(umkmSubCategories *[]UMKMSubCategory, umkmCategoryId string) (err error) {
	if err = Config.DB.Where("category_id = ?", umkmCategoryId).Find(umkmSubCategories).Error; err != nil {
		return err
	}
	return nil
}

func GetUMKMByLocationId(umkm *[]UMKM, locationId string) (err error) {
	if err = Config.DB.Where("location_id = ?", locationId).Find(umkm).Error; err != nil {
		return err
	}
	return nil
}

func GetFavoriteUMKM(umkm *[]UMKM) (err error) {
	if err = Config.DB.Order("score desc").Find(umkm).Error; err != nil {
		return err
	}
	return nil
}

func GetUMKMById(umkm *UMKM, umkmId string) (err error) {
	if err = Config.DB.Where("id = ?", umkmId).Find(umkm).Error; err != nil {
		return err
	}
	return nil
}

func GetUMKMSByOwnerId(umkm *[]UMKM, ownerId string) (err error) {
	if err = Config.DB.Where("owner_id = ?", ownerId).Find(umkm).Error; err != nil {
		return err
	}
	return nil
}

func GetUMKMSById(umkm *[]UMKM, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(umkm).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUMKM(umkm *UMKM) (err error) {
	Config.DB.Save(umkm)
	return nil
}

func GetUMKMSByName(umkms *[]UMKM, name string) (err error) {
	if err = Config.DB.Where("name like ?", "%"+name+"%").Find(umkms).Error; err != nil {
		return err
	}
	return nil
}

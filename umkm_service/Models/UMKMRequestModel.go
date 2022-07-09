package Models

type UMKMRequest struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`

	// Array datas
	FacilityIds    []string `json:"facility_ids"`
	Photos         []string `json:"photos"`
	SubCategoryIds []string `json:"subcategory_ids"`

	CategoryId int `json:"category_id"`
	ProvinceId int `json:"province_id"`
	LocationId int `json:"location_id"`
}

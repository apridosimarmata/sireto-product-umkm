package Models

type UMKM struct {
	Id             int    `json:"id"`
	Name           string `json:"name" gorm:"size:50"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	FacilityIds    string `json:"facility_ids"`
	Photos         string `json:"photos"`
	SubcategoryIds string `json:"subcategory_ids"`

	Score         float32 `json:"score"`
	ReviewsNumber int     `json:"reviews_number"`
	OwnerId       int     `json:"owner_id"`
	CategoryId    int     `json:"category_id"`
	ProvinceId    int     `json:"province_id"`
	LocationId    int     `json:"location_id"`

	Active bool `json:"active" gorm:"default:true"`
}

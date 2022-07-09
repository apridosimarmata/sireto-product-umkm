package Models

type UMKMCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	SubCategories []UMKMSubCategory `json:"subcategories"`
}

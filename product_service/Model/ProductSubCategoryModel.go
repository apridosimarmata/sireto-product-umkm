package model

type ProductSubCategory struct {
	Id int `json:"int"`

	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}

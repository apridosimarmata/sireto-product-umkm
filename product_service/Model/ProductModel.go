package model

import "time"

type Product struct {
	Id int `json:"id"`

	Sold         int `json:"sold"`
	Score        int `json:"score"`
	ReviewsCount int `json:"reviews_count"`
	CategoryId   int `json:"category_id"`
	MerchantId   int `json:"merchant_id"`
	Price        int `json:"price"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Photos      string `json:"photos"`

	CreatedAt time.Time `json:"created_at"`

	Halal  bool `json:"halal"`
	Active bool `json:"active"`
}

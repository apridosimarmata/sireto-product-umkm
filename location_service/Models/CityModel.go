package Models

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	ProvinceId int `json:"province_id"`
}

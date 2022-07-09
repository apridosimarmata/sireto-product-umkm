package Models

type Province struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	City []City `json:"cities"`
}

package Models

type Review struct {
	BusinessId     string   `json:"business_id"`
	Content        string   `json:"content"`
	Photos         []string `json:"photos"`
	PlaceScore     string   `json:"place_score"`
	ProductScore   string   `json:"product_score"`
	TreatmentScore string   `json:"treatment_score"`
}

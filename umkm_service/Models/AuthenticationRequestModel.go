package Models

type AuthenticationRequest struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

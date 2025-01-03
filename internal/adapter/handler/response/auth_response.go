package response

type SuccessAuthResponse struct {
	Meta
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}

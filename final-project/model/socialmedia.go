package model

type SocialMedia struct {
	SocialMediaID  int    `json:"social_media_id"`
	UserID         int    `json:"user_id"`
	SocialMediaURL string `json:"social_media_url"`
	Name           string `json:"name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

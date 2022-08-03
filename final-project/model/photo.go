package model

type Photo struct {
	PhotoID   int    `json:"photo_id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

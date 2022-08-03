package model

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentResponse struct {
	CommentID int       `json:"comment_id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		UserID   int    `json:"user_id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"User"`
	Photo struct {
		PhotoID  int    `json:"photo_id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoURL string `json:"photo_url"`
		UserID   int    `json:"user_id"`
	} `json:"Photo"`
}

package model

type User struct {
	UserID    int    `json:"id"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=4"`
	Age       int    `json:"age" validate:"required,min=0,max=100"`
	UpdatedAt string `json:"updated_at"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

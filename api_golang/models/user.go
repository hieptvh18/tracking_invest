package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Phone string `json:"phone"`
	Email string `json:"email"` // unique
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Age   string `json:"age" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Age   *string `json:"age"`
	Phone *string `json:"phone"`
	Email *string `json:"email"`
}
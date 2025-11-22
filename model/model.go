package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Role         string `json:"role" binding:"required"`
	Phone_Number int    `json:"phone_number" binding:"required"`
}

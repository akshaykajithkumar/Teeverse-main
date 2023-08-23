package models

import (
	"time"
)

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}

type TokenAdmin struct {
	Username string
	Token    string
}

type UserDetailsAtAdmin struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission bool   `json:"permission"`
}

// type AdminSalesReport struct{
// 	Orders	models.Order
// }

type CustomDates struct {
	StartingDate time.Time `json:"startingDate"`
	EndDate      time.Time `json:"endDate"`
}

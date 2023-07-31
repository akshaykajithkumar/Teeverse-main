package domain

type Admin struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Phone    string `gorm:"unique" json:"phone"`
	Password string `json:"password"`
	Active   bool   `gorm:"default:true" json:"active"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

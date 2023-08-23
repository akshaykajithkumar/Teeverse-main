package interfaces

import (
	"Teeverse/pkg/domain"
	"Teeverse/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserByID(id string) (domain.User, error)
	UpdateBlockUserByID(user domain.User) error
	GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error)
}

package interfaces

import (
	"Teeverse/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
	GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error)
}

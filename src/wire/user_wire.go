package wire

import (
	"service-user-management/src/controller"
	"service-user-management/src/repository"

	"gorm.io/gorm"
)

func UserWire(db *gorm.DB) controller.UserController {
	ur := repository.ProvideUserRepository(db)
	return controller.ProvideUserController(ur)
}

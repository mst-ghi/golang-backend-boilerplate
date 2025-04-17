package users

import (
	"app/database/db_scopes"
	"app/database/models"
	"app/database/repositories"
	"app/internal"
)

type UsersServiceInterface interface {
	FindOne(id string) (models.User, internal.Error)
	FindAll(search string, page, take int) ([]models.User, db_scopes.PaginateMetadata)
}

type UsersService struct {
	repository repositories.UserRepositoryInterface
}

func NewUsersService() *UsersService {
	return &UsersService{
		repository: repositories.NewUserRepository(),
	}
}

func (service *UsersService) FindOne(id string) (models.User, internal.Error) {
	user := service.repository.FindByID(id)

	if user.ID == "" {
		return user, internal.Error{"reason": "User not found"}
	}

	return user, nil
}

func (service *UsersService) FindAll(search string, page, take int) ([]models.User, db_scopes.PaginateMetadata) {
	return service.repository.FindAll(search, page, take)
}

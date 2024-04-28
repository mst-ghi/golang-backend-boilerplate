package repositories

import (
	"app/database"
	"app/database/db_scopes"
	"app/database/models"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Connection() *gorm.DB
	Create(user models.User) models.User
	FindByEmail(email string) models.User
	FindByID(id string) models.User
	FindAll(search string, page, take int) ([]models.User, db_scopes.PaginateMetadata)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (repo *UserRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	repo.DB.Create(&user).Scan(&newUser)
	return newUser
}

func (repo *UserRepository) FindByEmail(email string) models.User {
	var user models.User
	repo.DB.First(&user, "email = ?", email)
	return user
}

func (repo *UserRepository) FindByID(id string) models.User {
	var user models.User
	repo.DB.First(&user, "id = ?", id)
	return user
}

func (repo *UserRepository) FindAll(search string, page, take int) ([]models.User, db_scopes.PaginateMetadata) {
	var users []models.User

	query := repo.DB.
		Scopes(db_scopes.Paginate(page, take)).
		Table("users")

	if search != "" {
		query.Where("name LIKE ?", "%"+search+"%")
	}

	query.
		Order("created_at desc").
		Find(&users)

	var totalRows int64
	repo.DB.Table("users").Count(&totalRows)

	return users, db_scopes.PaginateMeta(totalRows, page, take)
}

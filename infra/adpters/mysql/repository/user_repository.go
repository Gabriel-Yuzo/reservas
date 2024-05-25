package repository

import (
	models "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint64) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func InitializeRepositories(db *gorm.DB) UserRepository {
	return NewUserRepository(db)
}

func (r *UserRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) GetUserByID(id uint64) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) DeleteUser(id uint64) error {
	return r.db.Delete(&models.User{ID: id}).Error
}

package repository

import (
	"github.com/Gabriel-Yuzo/reservas/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Save(User *models.User) error {
	return r.db.Create(User).Error
}

func (r *UserRepositoryImpl) FindByID(id int) (*models.User, error) {
	var User models.User
	err := r.db.First(&User, id).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (r *UserRepositoryImpl) FindAll() ([]*models.User, error) {
	var Users []*models.User
	err := r.db.Find(&Users).Error
	if err != nil {
		return nil, err
	}
	return Users, nil
}

func (r *UserRepositoryImpl) Update(User *models.User) error {
	return r.db.Save(User).Error
}

func (r *UserRepositoryImpl) Delete(id int) error {
	return r.db.Delete(id).Error
}

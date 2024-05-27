package usecase

import (
	"github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/models"
	"github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/repository"
)

type UserUsecase interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint64) (*models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) CreateUser(user *models.User) error {
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) GetUserByID(id uint64) (*models.User, error) {
	return u.userRepo.GetUserByID(id)
}

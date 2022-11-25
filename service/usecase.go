package service

import (
	"github.com/Fachrulmustofa20/go-microservice-user/models"
)

type UsersUsecase interface {
	Register(user models.Users) (userId uint64, err error)
	Login(email string, password string) (token string, err error)
	CreateProfile(profile models.Profile) (err error)
	UpdateProfileByUserId(profile models.Profile) (err error)
	GetUserById(userId uint64) (user models.Users, err error)
	GetProfileByUserId(userId uint64) (profile models.Profile, err error)
}

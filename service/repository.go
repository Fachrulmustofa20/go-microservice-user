package service

import (
	"github.com/Fachrulmustofa20/go-microservice-user/models"
)

type UsersRepository interface {
	CreateUser(user models.Users) (userId uint64, err error)
	GetUserByEmail(email string) (user models.Users, err error)
	GetUserById(userId uint64) (user models.Users, err error)
	CreateProfile(profile models.Profile) (err error)
	GetProfileByUserId(userId uint64) (profile models.Profile, err error)
	UpdateProfileByUserId(profile models.Profile) (err error)
}

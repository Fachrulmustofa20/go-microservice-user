package usecase

import (
	"fmt"

	"github.com/Fachrulmustofa20/go-microservice-user/constants"
	"github.com/Fachrulmustofa20/go-microservice-user/models"
	"github.com/Fachrulmustofa20/go-microservice-user/service"
	"github.com/Fachrulmustofa20/go-microservice-user/service/utils"
)

type usersUsecase struct {
	userRepo service.UsersRepository
}

func NewUsersUsecase(userRepo service.UsersRepository) service.UsersUsecase {
	return &usersUsecase{
		userRepo: userRepo,
	}
}

func (usecase usersUsecase) Register(users models.Users) (userId uint64, err error) {
	// hash password
	users.Password = utils.HashPass(users.Password)
	userId, err = usecase.userRepo.CreateUser(users)
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return userId, err
	}

	return userId, nil
}

func (usecase usersUsecase) Login(email string, password string) (token string, err error) {
	users, err := usecase.userRepo.GetUserByEmail(email)
	if err != nil {
		return token, constants.ErrLogin
	}

	comparePass := utils.ComparePassword([]byte(users.Password), []byte(password))
	if !comparePass {
		return token, constants.ErrLogin
	}

	token, err = utils.GenerateToken(users.ID, users.Email)
	if err != nil {
		return token, err
	}

	return token, nil
}

func (usecase usersUsecase) CreateProfile(profile models.Profile) (err error) {
	return usecase.userRepo.CreateProfile(profile)
}

func (usecase usersUsecase) UpdateProfileByUserId(profile models.Profile) (err error) {
	return usecase.userRepo.UpdateProfileByUserId(profile)
}

func (usecase usersUsecase) GetUserById(userId uint64) (user models.Users, err error) {
	return usecase.userRepo.GetUserById(userId)
}

func (usecase usersUsecase) GetProfileByUserId(userId uint64) (profile models.Profile, err error) {
	return usecase.userRepo.GetProfileByUserId(userId)
}

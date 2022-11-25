package postgres

import (
	"fmt"

	"github.com/Fachrulmustofa20/go-microservice-user/models"

	"gorm.io/gorm"
)

type usersRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usersRepository {
	return usersRepository{
		db: db,
	}
}

func (r usersRepository) CreateUser(user models.Users) (userId uint64, err error) {
	err = r.db.Create(&user).Error
	if err != nil {
		err = fmt.Errorf("%+v", err)
		return user.ID, err
	}
	return user.ID, nil
}

func (r usersRepository) GetUserById(userId uint64) (user models.Users, err error) {
	err = r.db.Debug().Where("id = ?", userId).Take(&user).Error
	if err != nil {
		fmt.Printf("[UserRepository][GetUserByEmail] error while check user by email: %+v", err)
		return user, err
	}
	return user, err
}

func (r usersRepository) GetUserByEmail(email string) (user models.Users, err error) {
	err = r.db.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		fmt.Printf("[UserRepository][GetUserByEmail] error while check user by email: %+v", err)
		return user, err
	}
	return user, err
}

func (r usersRepository) UpdateProfileByUserId(profile models.Profile) (err error) {
	if err = r.db.Debug().Model(profile).Where("user_id = ?", profile.UserId).Updates(models.Profile{
		Age:         profile.Age,
		Photo:       profile.Photo,
		Hoby:        profile.Hoby,
		Description: profile.Description,
	}).Error; err != nil {
		fmt.Printf("[UserRepository][UpdateProfileByUserId] error while update profile by user_id: %+v", err)
		return err
	}

	return nil
}

func (r usersRepository) CreateProfile(profile models.Profile) (err error) {
	err = r.db.Debug().Create(&profile).Error
	if err != nil {
		fmt.Printf("[usersRepository][UpdateProfile] error while update profile: %+v", err)
		return err
	}
	return nil
}

func (r usersRepository) GetProfileByUserId(userId uint64) (profile models.Profile, err error) {
	err = r.db.Debug().Where("user_id = ?", userId).First(&profile).Error
	if err != nil {
		fmt.Printf("[usersRepository][GetProfileByUserId] error while get profile by user id : %+v", err)
		return profile, err
	}
	return profile, err
}

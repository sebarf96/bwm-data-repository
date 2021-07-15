package user_service

import (
	"bwm-api-repository/src/models"
	"bwm-api-repository/src/repositories/user_repository"
	"github.com/go-kit/kit/log"
)

type UserService interface {
	Create(user models.User) (string,error)
	Read() (models.Users, error)
	Update(user models.User, id string) error
	Delete(id string) error
}

type userService struct {
	log            log.Logger
	userRepository user_repository.UserRepository
}

func NewUserService(log log.Logger, repository user_repository.UserRepository) UserService {

	return &userService{
		log:            log,
		userRepository: repository,
	}
}

func (u *userService) Create(user models.User) (string,error) {

	err := u.userRepository.Create(user)

	if err != nil {
		return "",err
	}

	return "User created succesfully",nil
}


func (u *userService) Read() (models.Users, error) {
	users, err := u.userRepository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) Update(user models.User, id string) error {

	err := u.userRepository.Update(user, id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Delete(id string) error {

	err := u.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

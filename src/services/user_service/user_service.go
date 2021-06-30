package user_service

import (
	"bwm-api-repository/src/models"
	"bwm-api-repository/src/repositories/user_repository"
)

func Create(user models.User) error {

	err := user_repository.Create(user)

	if err != nil {
		return err
	}

	return nil
}

func Read() (models.Users, error) {
	users, err := user_repository.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Update(user models.User, id string) error {

	err := user_repository.Update(user, id)

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {

	err := user_repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

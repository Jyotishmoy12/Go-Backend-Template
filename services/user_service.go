package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserById() error
}

type UserServiceImpl struct {
	UserRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("fetching user in user service")
	users, err :=u.UserRepository.GetAll()

	if err != nil {
		fmt.Println("Error fetching users", err)
		return err
	}
	fmt.Println("Users fetched successfully")

	for _, user:= range users {
		fmt.Println(user)
	}
	return nil
}

// structs can implement a particular interface how can we make sure that the struct implements the interface?
// by implementing all the functions of the interface it will work
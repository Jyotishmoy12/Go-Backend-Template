package db

import (
	// "database/sql"
	"fmt"
)

// interface for user
// dependency injection is a design pattern that allows a class to depend on other classes without having to know their concrete implementations
type UserRepository interface {
	Create() error
}

// implementation for user Implementation of struct
type UserRepositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u * UserRepositoryImpl) Create() error {
	fmt.Println("Creating user in user repository")
	return nil
}
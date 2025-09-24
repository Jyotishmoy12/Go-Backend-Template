package db

import (
	// "database/sql"
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// interface for user
// dependency injection is a design pattern that allows a class to depend on other classes without having to know their concrete implementations
type UserRepository interface {
	GetByID() (*models.User, error) // return user and error
	Create() error
}

// implementation for user Implementation of struct
type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("creating user in user repository")

	// step 1: prepare the query
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	// step 2: execute the query
	// Exec function is used to execute a query that doesn't return any rows
	result, err := u.db.Exec(query, "testuser", "test@test.com", "password123")

	if err != nil {
		fmt.Println("Error creating user", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()

	if rowsAffected == 0 {
		fmt.Println("No rows were affected user not created", rowErr)
		return nil
	}

	fmt.Println("User created successfully", rowsAffected)

	return nil
}

func (u *UserRepositoryImpl) GetByID() (*models.User, error) { // return user and error {
	fmt.Println("fetching user in user repository")

	// step 1: prepare the query
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

	// step 2: execute the query
	// queryRow function is used to execute a query that returns a single row
	row := u.db.QueryRow(query, 1)

	// step 3: process this row and create output object
	user := &models.User{}

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with id 1")
			return nil, err
		} else {
			fmt.Println("Error fetching user", err)
			return nil, err
		}
	}
	// step 4: print user details
	fmt.Println("User fetched successfully", user)

	return user, nil
}

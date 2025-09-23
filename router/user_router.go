package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

// here no dependency injection because we are using controllers and controllers doesnot have business logic in it
type UserRouter struct {
	UserController *controllers.UserController
}


func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/signup", ur.UserController.RegisterUser)
}
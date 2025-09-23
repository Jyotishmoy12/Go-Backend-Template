package router

import (
	"AuthInGo/controllers"
	"github.com/go-chi/chi/v5"
)

// Router interface 
// i will defind the router only when it implements this interface
type Router interface {
	Register(r chi.Router)
}

func SetUpRouter(UserRouter Router) * chi.Mux {

	ChiRouter := chi.NewRouter()

	ChiRouter.Get("/ping", controllers.PingHandler)

	UserRouter.Register(ChiRouter)
   
	return ChiRouter

}
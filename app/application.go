package app

import (
	dbConfig "AuthInGo/config/db"
	"AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // PORT
}


// constructor for config
func NewConfig() Config {
	port := config.GetString("PORT", ":8080")
	return Config{
		Addr: port,
	}
}
// application struct which has config struct
// it will get all the configuration from config struct
type Application struct {
	Config Config
	Store db.Storage
}

// constructor for application
func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
		Store: *db.NewStorage(),
	}
}
// member function of application struct
// it will run the server
// it will take the Application instance 
func (app *Application) Run() error {
	db, err :=dbConfig.SetupDB()
	if err != nil {
		fmt.Println("Failed to connect to the database", err)
		return err
	}
	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetUpRouter(uRouter),
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
	}
	fmt.Println("Starting server on", app.Config.Addr)
	return server.ListenAndServe()
}
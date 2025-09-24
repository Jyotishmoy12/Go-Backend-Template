package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
	"fmt"
)

func main() {
    
	config.Load()
	
	cfg := app.NewConfig()
	// so now wee need an application object to call the struct
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()

	app.Run()
	fmt.Println("Server is running on", cfg.Addr)
}

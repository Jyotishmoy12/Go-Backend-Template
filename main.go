package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	"fmt"
)

func main() {
    
	config.Load()
	
	cfg := app.NewConfig()
	// so now wee need an application object to call the struct
	app := app.NewApplication(cfg)


	app.Run()
	fmt.Println("Server is running on", cfg.Addr)
}

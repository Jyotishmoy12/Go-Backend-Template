package config

import (
	"database/sql"
	"fmt"

	env "AuthInGo/config/env"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "localhost:3306")
	cfg.DBName = env.GetString("DB_NAME", "auth_in_go")

	fmt.Println("Connecting to database", cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Failed to connect to the database", err)
		return nil, err
	}
	fmt.Println("Trying to connect to the db")
	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Failed to connect to the database", pingErr)
		return nil, pingErr
	}
	fmt.Println("Database connected successfully")

	return db, nil
}

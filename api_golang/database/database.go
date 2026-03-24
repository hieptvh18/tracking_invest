package database

import (
	"api_golang/utils/env"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectMySQL returns *sql.DB for MySQL using env: DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME.
func ConnectMySQL() (*sql.DB, error) {
	host := env.GetEnv("DB_HOST", "localhost")
	port := env.GetEnv("DB_PORT", "3306")
	user := env.GetEnv("DB_USER", "root")
	pass := env.GetEnv("DB_PASSWORD", "")
	name := env.GetEnv("DB_NAME", "golang")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("MySQL connected: %s:%s/%s", host, port, name)
	return db, nil
}

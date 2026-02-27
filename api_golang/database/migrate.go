package database

import (
	"database/sql"
	"log"
)

const usersTableSQL = `
CREATE TABLE IF NOT EXISTS users (
	id    VARCHAR(36) PRIMARY KEY AUTO_INCREMENT,
	name  VARCHAR(255) NOT NULL,
	age   VARCHAR(50)  NOT NULL,
	phone VARCHAR(50)  NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE
);
`

// MigrateUsers creates users table if not exists.
func MigrateUsers(db *sql.DB) error {
	_, err := db.Exec(usersTableSQL)
	if err != nil {
		return err
	}
	log.Println("Migration: users table ready")
	return nil
}

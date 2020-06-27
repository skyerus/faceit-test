package db

import (
	"database/sql"
	"os"
)

// OpenDb - create db connection
func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_URL")+"?parseTime=true")
	if err != nil {
		return db, err
	}
	return db, nil
}

// InitiateTables - create database tables if don't already exist
func InitiateTables(conn *sql.DB) error {
	return createUserTable(conn)
}

func createUserTable(conn *sql.DB) error {
	_, err := conn.Exec("CREATE TABLE IF NOT EXISTS `faceit`.`user` (`id` INT NOT NULL AUTO_INCREMENT, `first_name` VARCHAR(45) NULL,`last_name` VARCHAR(45) NULL, `nickname` VARCHAR(45) NULL, `email` VARCHAR(45) NULL, `country` VARCHAR(45) NULL, PRIMARY KEY (`id`), UNIQUE INDEX `nickname_UNIQUE` (`nickname` ASC), UNIQUE INDEX `email_UNIQUE` (`email` ASC), INDEX `index4` (`country` ASC));")

	return err
}

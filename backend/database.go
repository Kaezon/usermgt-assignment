package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DatabaseConnect(c config) (*sql.DB, error) {
	db, err := sql.Open(
		c.DatabaseDriver,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			c.DatabaseHost,
			c.DatabasePort,
			c.DatabaseUsername,
			c.DatabasePassword,
			c.DatabaseName,
		),
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping();
	err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeDatabase(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE users (
	  user_id BIGSERIAL PRIMARY KEY,
	  user_name VARCHAR(50) NOT NULL,
	  first_name VARCHAR(255) NOT NULL,
	  last_name VARCHAR(255) NOT NULL,
	  email VARCHAR(255) NOT NULL,
	  user_status VARCHAR(1) NOT NULL,
	  department VARCHAR(255),
	);
	`)
	if err != nil {
		return err
	}
	return nil
}
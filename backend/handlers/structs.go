package handlers

import "database/sql"

type DatabaseHandler struct {
	DB *sql.DB
}

type UsersHandler DatabaseHandler

func NewUsersHandler(db *sql.DB) *UsersHandler {
	return &UsersHandler{DB: db}
}
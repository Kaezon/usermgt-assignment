package models

import "github.com/gocraft/dbr"

type User struct {
	ID			int64
	Username 	string
	FirstName	string
	LastName	string
	Email		string
	UserStatus	string
	Department	dbr.NullString
}
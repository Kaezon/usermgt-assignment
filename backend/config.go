package main

import "os"

type config struct {
	DatabaseDriver		string
	DatabaseName		string
	DatabaseHost		string
	DatabasePort		string
	DatabaseUsername	string
	DatabasePassword	string
}

func GetConfig() config {
	return config{
		DatabaseDriver: os.Getenv("DB_DRIVER"),
		DatabaseName: os.Getenv("DB_NAME"),
		DatabaseHost: os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
		DatabaseUsername: os.Getenv("DB_USERNAME"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
	}
}
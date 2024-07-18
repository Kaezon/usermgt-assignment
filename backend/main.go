package main

import (
	"log"
	"os"

	"usermgt/handlers"

	"github.com/labstack/echo"
)

func main() {
	conf := GetConfig()
	e := echo.New()

	db, err := DatabaseConnect(conf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	userHandler := handlers.NewUsersHandler(db)

	// Register routes
	e.DELETE("/users/:id", userHandler.DeleteUser)
	e.GET("/users", userHandler.ListUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)

	log.Fatal(e.Start(":8080"))
}

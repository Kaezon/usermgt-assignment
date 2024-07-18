package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"usermgt/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/labstack/echo"
)

func (h *UsersHandler) CreateUser(c echo.Context) error {
	var user models.User
	var userID int64
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

	if err := sq.
		Insert("users").
		Columns("user_name",
				"first_name",
				"last_name",
				"email",
				"user_status",
				"department").
		Values(user.Username,
			   user.FirstName,
			   user.LastName,
			   user.Email,
			   user.UserStatus,
			   user.Department).
		Suffix("RETURNING user_id").
		RunWith(h.DB).
		PlaceholderFormat(sq.Dollar).
		QueryRow().
		Scan(&userID);
	err != nil {
		log.Println("Query error: ", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	} else {
		return c.JSON(http.StatusCreated, userID)
	}
}

func (h *UsersHandler) DeleteUser(c echo.Context) error {
	if result, err := sq.
		Delete("users").
		Where(sq.Eq{"user_id":c.Param("id")}).
		RunWith(h.DB).
		PlaceholderFormat(sq.Dollar).
		Exec();
	err != nil {
		log.Println("Query error: ", err)
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	} else if count, _ := result.RowsAffected(); count < 1 {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, nil)
}

func (h *UsersHandler) GetUser(c echo.Context) error {
	var user models.User

	if err := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"user_id": c.Param("id")}).
		RunWith(h.DB).
		PlaceholderFormat(sq.Dollar).
		QueryRow().
		Scan(&user.ID,
			 &user.Username,
			 &user.FirstName,
			 &user.LastName,
			 &user.Email,
			 &user.UserStatus,
			 &user.Department);
	err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "User not found")
		} else {
			log.Println("Query error: ", err)
			return c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UsersHandler) ListUsers(c echo.Context) error {
	rows, err := sq.
		Select("*").
		From("users").
		RunWith(h.DB).
		Query()
	if err != nil {
		log.Println("Query error: ", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID,
							&user.Username,
							&user.FirstName,
							&user.LastName,
							&user.Email,
							&user.UserStatus,
							&user.Department);
		err != nil {
			log.Println("Query error: ", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		users = append(users, user)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UsersHandler) UpdateUser(c echo.Context) error {
    var user models.User
	var userID int64

    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

	query := sq.
		Update("users").
		Where(sq.Eq{"user_id":c.Param("id")}).
		Suffix("RETURNING user_id")
	
	if user.Username != "" {
		query = query.Set("user_name", user.Username)
	}
	if user.FirstName != "" {
		query = query.Set("first_name", user.FirstName)
	}
	if user.LastName != "" {
		query = query.Set("last_name", user.LastName)
	}
	if user.Email != "" {
		query = query.Set("email", user.Email)
	}
	if user.UserStatus != "" {
		query = query.Set("user_status", user.UserStatus)
	}
	if user.Department.Valid {
		query = query.Set("department", user.Department.String)
	}

	if err := query.
		RunWith(h.DB).
		PlaceholderFormat(sq.Dollar).
		QueryRow().
		Scan(&userID)
	err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "User not found")
		} else {
			log.Println("Query error: ", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, userID)
}
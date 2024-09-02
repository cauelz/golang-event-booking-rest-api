package models

import (
	"errors"

	"github.com/cauelz/golang-event-booking-rest-api/db"
	"github.com/cauelz/golang-event-booking-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {

	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, error := db.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer stmt.Close()

	hashedPassword, error := utils.HashPassword(u.Password)

	if error != nil {
		return error
	}

	result, error := stmt.Exec(u.Email, hashedPassword)

	if error != nil {
		return error
	}

	id, error := result.LastInsertId()

	u.ID = id

	return error
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	error := row.Scan(&u.ID, &retrievedPassword)

	if error != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
package models

import (
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

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	
	// GenerateFromPassword returns the bcrypt hash of the password at the given cost.
	// Cost means the cost of the hashing function. The higher the cost, the more secure the hash.
	// The bcrypt package uses the cost of 10 by default.
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), error
}
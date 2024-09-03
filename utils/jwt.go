package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "my_secret_key"

func GenerateToken(email string, userId int64) (string, error) {

	// Create a new token object, specifying signing method and the claims
	// signing method: jwt.SigningMethodHS256 which means that the token will be signed using the HMAC-SHA algorithm
	// claims: jwt.MapClaims which is a map[string]interface{} that will be used to store the claims.
	// The claims are the data that will be stored in the token. In this case, we are storing the email, userId, and the expiration time of the token.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret key
	// secret key means that the token will be signed using the HMAC-SHA algorithm with the secret key as the secret.
	return jwtToken.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {

	// Parse the token using the secret key
	parsedToken, error := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if error != nil {
		return errors.New("Could not verify token")
	}

	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}

	// claims, ok :=parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("Could not parse claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(float64)

	return nil
}
package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Create a token by id

func GenToken(userID int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))

}

// valid token extract from request by parameter authorization header

func ValidToken(r *http.Request) error {
	extractTknString := extractToken(r)
	token, err := jwt.Parse(extractTknString, returnKey)
	if err != nil {
		return err
	}

	// validate the token by jwt.claims from lib
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalid")
}

// extract token from request from parameter authorization and split the returned string. because return has 1[bearer token] 2[token]
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// returning the secret key from .inv
func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo invalid %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractIDfromToken(r *http.Request) (uint64, error) {
	extractTknString := extractToken(r)

	token, err := jwt.Parse(extractTknString, returnKey)

	if err != nil {
		return 0, err
	}
	fmt.Println("token aqui no token.go", token)

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permission["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, errors.New("tokeninvalid")

}

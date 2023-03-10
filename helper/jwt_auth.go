package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"my-social-network/config"

	jwt "github.com/golang-jwt/jwt/v4"
)

// CreateJwt returns a token with all the user's permissions
func CreateJwt(userID string) (string, error) {

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["user_id"] = userID
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	fmt.Printf("userID: %v\n", userID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, returnAuthKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	barearToken := strings.Split(token, " ")

	if len(barearToken) == 2 {
		return barearToken[1]
	}
	return ""
}

func returnAuthKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

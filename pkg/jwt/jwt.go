package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(id string) (string, error) {

	claim := jwt.MapClaims{
		"sum": id,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
		"iat": time.Now().Unix(),	
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))

	if err != nil {
		return "", err
	}

	return t, nil

}

func ValidateToken(token string) bool {

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("token invalido")
		}

		return []byte(os.Getenv("JWTSECRET")), nil
	})

	return err == nil

}
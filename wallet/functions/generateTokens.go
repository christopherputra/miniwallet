package functions

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var secretKey = "8UYAJSGFU124G"

func GenerateToken(walletId string) (string, error) {
    claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"wallet_id": walletId ,
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims)
	
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
 	return tokenString, nil
}
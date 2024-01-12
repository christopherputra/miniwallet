package functions
import (
	"github.com/golang-jwt/jwt"
)
func VerifyToken(tokenString string) (valid bool, walletId string, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	
	valid = token.Valid
	if !valid {
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]interface{})
	walletId = data["wallet_id"].(string)

	return
}
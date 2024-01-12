package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	funcs "wallet/functions"
	"clevergo.tech/jsend"
)

func TokenAuthenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.GetHeader("Authorization")
		token := auth[len("Token "):]
		valid, walletId, err := funcs.VerifyToken(token)
		if auth == "" {
			c.IndentedJSON(http.StatusUnauthorized, jsend.NewFail(map[string]string{
				"error": "Unauthorized.",
			}))
			c.Abort()
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, jsend.NewFail(map[string]string{
				"error": "Unauthorized.",
			}))
			c.Abort()
			return
		}
		if !valid {
			c.IndentedJSON(http.StatusUnauthorized, jsend.NewFail(map[string]string{
				"error": "Unauthorized.",
			}))
			c.Abort()
			return
		}

		c.Set("wallet_id", walletId)
		c.Next()
	  }
}
package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"clevergo.tech/jsend"
	"wallet/service"
)

func WalletCheckingMiddleware(walletService *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		walletId := c.GetString("wallet_id")
		req := service.GetStatusReq{
			WalletId: walletId,
		}
		res := &service.GetStatusRes{}
		err := walletService.GetStatus(req, res)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail([]string{}))
			c.Abort()
			return
		}
		if res.Status == "disabled" {
			c.IndentedJSON(http.StatusBadRequest, jsend.NewFail(map[string]string{
				"error": "Wallet Disabled.",
			}))
			c.Abort()
			return
		}

		c.Next()
	  }
}
package handler
import (
	"fmt"
	"wallet/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"clevergo.tech/jsend"
	"github.com/go-playground/validator/v10"
	"strings"
	"reflect"
)

type Handler struct {
	WalletService service.Service
}

func parseError(err error, blueprint interface{}) map[string][]string {
	out := make(map[string][]string)
		switch typedError := any(err).(type) {
		case validator.ValidationErrors:
			for _, e := range typedError {
				fieldName := e.Field()
				field, _ := reflect.TypeOf(blueprint).Elem().FieldByName(fieldName)
				fieldJSONName, _ := field.Tag.Lookup("json")
				out[fieldJSONName] = parseFieldError(e)
			}
	}
	
	return out
}

func parseFieldError(e validator.FieldError) []string {
	tag := strings.Split(e.Tag(), "|")[0]
	switch tag {
	case "required":
		return []string{"Missing data for required field."}
	default:
		return []string{"Something is not right."}
	}
}


func (h *Handler) InitWallet(c *gin.Context) {
	var req service.InitWalletReq
	res := &service.InitWalletRes{}
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsend.NewFail(parseError(err, &req)))
        return
    }

	err := h.WalletService.InitWallet(req,res)

	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}

func (h *Handler) EnableWallet(c *gin.Context) {
	var req service.EnableWalletReq
	res := &service.EnableWalletRes{}
	
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.EnableWallet(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}

func (h *Handler) DisableWallet(c *gin.Context) {
	var req service.DisableWalletReq
	res := &service.DisableWalletRes{}
	
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.DisableWallet(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}


func (h *Handler) GetWallet(c *gin.Context) {
	var req service.GetWalletReq
	res := &service.GetWalletRes{}
	
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.GetWallet(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}


func (h *Handler) DepositTransaction(c *gin.Context) {
	var req service.DepositTransactionReq
	res := &service.DepositTransactionRes{}
	
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsend.NewFail(parseError(err, &req)))
        return
    }
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.DepositTransaction(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}

func (h *Handler) WithdrawTransaction(c *gin.Context) {
	var req service.WithdrawTransactionReq
	res := &service.WithdrawTransactionRes{}
	
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, jsend.NewFail(parseError(err, &req)))
        return
    }
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.WithdrawTransaction(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}

func (h *Handler) ListTransactions(c *gin.Context) {
	var req service.ListTransactionsReq
	res := &service.ListTransactionsRes{}
	
	req.WalletId = c.GetString("wallet_id")

	err := h.WalletService.ListTransactions(req,res)
	if err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, jsend.NewFail(map[string]string{
			"error": err.Error(),
		}))
		c.Abort()
		return
	}
	
    c.IndentedJSON(http.StatusCreated, jsend.New(res))
	return
}



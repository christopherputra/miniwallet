package main

import (
	"context"
	"wallet/database"
	"fmt"
	"io/ioutil"
	"encoding/json"
	service "wallet/service"
	"github.com/gin-gonic/gin"
	"log"
	h "wallet/service/handler"
	middlewares "wallet/middlewares"
)

type Config struct {
	Postgres struct {
		Host string `json:"host"`
		Port int `json:"port"`
		User string `json:"user"`
		Password string `json:"password"`
		Dbname string `json:"dbname"`
	} `json:"postgres"`
    Host string `json:"host"`
    Port string `json:"port"`
}

func LoadConfig(file string) (Config, error) {
    var config Config
    configFile, err := ioutil.ReadFile(file)
    if err != nil {
        return config, err
    }
    err = json.Unmarshal(configFile, &config)
    if err != nil {
        return config, err
    }
    return config, nil
}

func main() {
	
	//LOAD CONFIG
	config, err := LoadConfig("./config.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	//POSTGRES
	postgresClient := database.NewPostgresClient(
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Dbname,
	);
	walletService := service.NewWalletService(postgresClient)

	handler := h.Handler {
		WalletService: walletService,
	}
	router := gin.New()
	router.POST("/api/v1/init", handler.InitWallet)

	rg1 := router.Group("/api/v1/wallet", middlewares.TokenAuthenticateMiddleware())
	rg1.POST("/", handler.EnableWallet)

	rg2 := router.Group("/api/v1/wallet", middlewares.TokenAuthenticateMiddleware(), middlewares.WalletCheckingMiddleware(&walletService))
	rg2.GET("/", handler.GetWallet)
	rg2.PATCH("/", handler.DisableWallet)
	rg2.POST("/deposits", handler.DepositTransaction)
	rg2.POST("/withdrawals", handler.WithdrawTransaction)
	rg2.GET("/transactions", handler.ListTransactions)
	if err := router.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		log.Fatal(context.Background(), err.Error())
	}
	

}

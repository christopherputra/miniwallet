package service;

import(
	"wallet/database"
)

type Wallet struct {
	WalletId string `json:"id"`
	OwnerId string `json:"owned_by"`
	Status string `json:"status"`
	Balance int64 `json:"balance"`
}

type Service struct {
	PgServer database.PostgresClient
}

func NewWalletService(pgserver database.PostgresClient) Service{
	return Service{
		PgServer: pgserver,
	}
}

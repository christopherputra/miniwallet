package service

import (
	"wallet/database"
)

type CalculateBalanceReq struct {
	WalletId string
}

type CalculateBalanceRes struct {
	Balance int64
}

func (w *Service) CalculateBalance(req CalculateBalanceReq, res *CalculateBalanceRes) error {

	reqListDb := database.ListTransactionsReqDb {
		WalletId: req.WalletId,
	}

	resTransDb, err := w.PgServer.ListTransactions(reqListDb)
	if err != nil {
		return err
	}

	var balance int64
	for _, transDb := range resTransDb.TransactionDbs {
		if transDb.Type == "deposit" {
			balance = balance + transDb.Amount
		} else {
			balance = balance - transDb.Amount
		}
	}
	res.Balance = balance

	return nil
}
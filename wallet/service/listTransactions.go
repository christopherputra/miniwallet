package service

import (
	"wallet/database"
	funcs "wallet/functions"
)

type ListTransactionsReq struct {
	WalletId string
}

type TransactionData struct {
	Transaction
	Type string `json:"type"`
	TransactedAt string `json:"transacted_at"`

}

type ListTransactionsRes struct {
	Transactions []TransactionData `json:"transactions"`
}

func (w *Service) ListTransactions(req ListTransactionsReq, res *ListTransactionsRes) error {

	reqDb := database.ListTransactionsReqDb {
		WalletId: req.WalletId,
	}

	resDb, err := w.PgServer.ListTransactions(reqDb)
	if err != nil {
		return err
	}

	transactions := make([]TransactionData, 0)
	for _, transDb := range resDb.TransactionDbs {
		transaction := TransactionData {
			Transaction: Transaction{
				TransactionId: transDb.TransactionId,
				Status: "success",
				Amount: transDb.Amount,
				ReferenceId: transDb.ReferenceId,
			},
			Type: transDb.Type,
			TransactedAt: funcs.UnixToTimeStamp(transDb.ExecutedAt),
		}
		transactions = append(transactions, transaction)
	}
	*res = ListTransactionsRes{
		Transactions: transactions,
	}
	return nil
}
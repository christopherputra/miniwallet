package service

import (
	"wallet/database"
	//funcs "wallet/functions"
	"time"
	"github.com/google/uuid"
	funcs "wallet/functions"
)

type WithdrawTransactionReq struct {
	WalletId string
	Amount int64 `json:"amount" binding:"required"`
	ReferenceId string `json:"reference_id" binding:"required"`
}

type WithdrawTransactionRes struct {
	WalletId string `json:"withdrawn_by"`
	Transaction
	WithdrawnAt string `json:"withdrawn_at"`
}

func (w *Service) WithdrawTransaction(req WithdrawTransactionReq, res *WithdrawTransactionRes) error {
	newUUID := uuid.NewString()
	reqDb := database.InsertTransactionReqDb {
		TransactionId: newUUID,
		Type: "withdrawal",
		WalletId: req.WalletId,
		Amount: req.Amount,
		ReferenceId: req.ReferenceId,
		ExecutedAt: time.Now().Unix(),
	}

	resDb, err := w.PgServer.InsertTransaction(reqDb)
	if err != nil {
		return err
	}

	transactionData := Transaction {
		TransactionId: resDb.TransactionId,
		Status: "success",
		Amount: resDb.Amount,
		ReferenceId: resDb.ReferenceId,
	}
	*res = WithdrawTransactionRes{
		Transaction: transactionData,
		WalletId: resDb.WalletId,
		WithdrawnAt: funcs.UnixToTimeStamp(resDb.ExecutedAt),
	}
	return nil
}
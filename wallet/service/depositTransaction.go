package service

import (
	"wallet/database"
	//funcs "wallet/functions"
	"time"
	"github.com/google/uuid"
	funcs "wallet/functions"
)

type DepositTransactionReq struct {
	WalletId string
	Amount int64 `json:"amount" binding:"required"`
	ReferenceId string `json:"reference_id" binding:"required"`
}

type Transaction struct {
	TransactionId string `json:"id"`
	Status string `json:"status"`
	Amount int64 `json:"amount"`
	ReferenceId string `json:"reference_id"`
}

type DepositTransactionRes struct {
	WalletId string `json:"deposited_by"`
	Transaction
	DepositedAt string `json:"deposited_at"`
}

func (w *Service) DepositTransaction(req DepositTransactionReq, res *DepositTransactionRes) error {
	newUUID := uuid.NewString()
	reqDb := database.InsertTransactionReqDb {
		TransactionId: newUUID,
		Type: "deposit",
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
	*res = DepositTransactionRes{
		Transaction: transactionData,
		WalletId: resDb.WalletId,
		DepositedAt: funcs.UnixToTimeStamp(resDb.ExecutedAt),
	}
	return nil
}
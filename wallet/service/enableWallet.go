package service

import (
	"wallet/database"
	funcs "wallet/functions"
	"time"
)

type EnableWalletReq struct {
	WalletId string
}

type EnabledWallet struct {
	Wallet
	EnabledAt string `json:"enabled_at"`
}

type EnableWalletRes struct {
	WalletData EnabledWallet `json:"wallet"`
}

func (w *Service) EnableWallet(req EnableWalletReq, res *EnableWalletRes) error {

	reqDb := database.EnableWalletReqDb {
		WalletId: req.WalletId,
		Status: "enabled",
		EnabledAt: time.Now().Unix(),
	}

	resDb, err := w.PgServer.EnableWallet(reqDb)
	if err != nil {
		return err
	}

	reqBalance := CalculateBalanceReq {
		WalletId: req.WalletId,
	}
	resBalance := &CalculateBalanceRes{}
	err = w.CalculateBalance(reqBalance, resBalance)
	if err != nil {
		return err
	}

	*res = EnableWalletRes{
		WalletData: EnabledWallet{
			Wallet: Wallet{
				WalletId: resDb.WalletId,
				OwnerId: resDb.CustomerId,
				Status: resDb.Status,
				Balance: resBalance.Balance,
			},
			EnabledAt: funcs.UnixToTimeStamp(resDb.EnabledAt),
		},
	}
	return nil
}
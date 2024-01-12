package service

import (
	"wallet/database"
	funcs "wallet/functions"
	"time"
)

type DisableWalletReq struct {
	WalletId string
}

type DisabledWallet struct {
	Wallet
	DisabledAt string `json:"disabled_at"`
}

type DisableWalletRes struct {
	WalletData DisabledWallet `json:"wallet"`
}

func getStatusName(status bool) string {
	if status {
		return "enabled"
	}
	return "disabled"
}

func (w *Service) DisableWallet(req DisableWalletReq, res *DisableWalletRes) error {

	reqDb := database.DisableWalletReqDb {
		WalletId: req.WalletId,
		Status: "disabled",
		DisabledAt: time.Now().Unix(),
	}

	resDb, err := w.PgServer.DisableWallet(reqDb)
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

	*res = DisableWalletRes{
		WalletData: DisabledWallet{
			Wallet: Wallet {
				WalletId: resDb.WalletId,
				OwnerId: resDb.CustomerId,
				Status: resDb.Status,
				Balance: resBalance.Balance,
			},
			DisabledAt: funcs.UnixToTimeStamp(resDb.DisabledAt),
		},
	}
	return nil
}
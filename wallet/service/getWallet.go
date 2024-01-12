package service

import (
	"wallet/database"
	funcs "wallet/functions"
)

type GetWalletReq struct {
	WalletId string
}

type GetWalletRes struct {
	WalletData EnabledWallet `json:"wallet"`
}

func (w *Service) GetWallet(req GetWalletReq, res *GetWalletRes) error {

	reqDb := database.GetWalletReqDb {
		WalletId: req.WalletId,
	}

	resDb, err := w.PgServer.GetWallet(reqDb)
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

	*res = GetWalletRes{
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
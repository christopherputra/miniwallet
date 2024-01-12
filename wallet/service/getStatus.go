package service

import (
	"wallet/database"
)

type GetStatusReq struct {
	WalletId string
}

type GetStatusRes struct {
	Status string
}

func (w *Service) GetStatus(req GetStatusReq, res *GetStatusRes) error {

	reqDb := database.GetWalletReqDb {
		WalletId: req.WalletId,
	}

	resDb, err := w.PgServer.GetWallet(reqDb)
	if err != nil {
		return err
	}

	*res = GetStatusRes{
		Status: resDb.Status,
	}
	return nil
}
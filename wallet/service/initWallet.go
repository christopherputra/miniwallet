package service

import (
    "github.com/google/uuid"
	"wallet/database"
	funcs "wallet/functions"
)

type InitWalletReq struct {
	CustomerId string `json:"customer_xid" binding:"required"`
}

type InitWalletRes struct {
	Token string `json:"token"`
}

func (w *Service) InitWallet(req InitWalletReq, res *InitWalletRes) error {
	newUUID := uuid.NewString()
	reqDb := database.CreateWalletReqDb {
		WalletId: newUUID,
		CustomerId: req.CustomerId,
		Status: "disabled",
	}
	resDb, err := w.PgServer.CreateOrInsertWallet(reqDb)
	if err != nil {
		return err
	}
	token, err := funcs.GenerateToken(resDb.WalletId)
	if err != nil {
		return err
	}
	*res = InitWalletRes{
		Token: token,
	}
	return nil
}
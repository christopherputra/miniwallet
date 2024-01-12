package service

import (
	"testing"
	mock_database "wallet/database/mock_database"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"wallet/database"
	"errors"
)

func TestGetStatusWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
    postgresClientMock := mock_database.NewMockPostgresClient(mockCtrl)
	errSample := errors.New("Error")

	
	t.Run("When postgres client is running properly", func(t *testing.T) {
		resDb := database.GetWalletResDb{
			WalletId: "wallet-id-1",
			CustomerId: "customer-id-1",
			Status: "enabled",
			EnabledAt: 18272491237,
		}
		postgresClientMock.EXPECT().GetWallet(gomock.Any()).Do(func (reqDb database.GetWalletReqDb) {
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(resDb, nil)

		req := GetStatusReq{
			WalletId: "wallet-id-1",
		}
		res := &GetStatusRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.GetStatus(req, res)

		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, res.Status, "enabled", "Status is not the same")

	})

	t.Run("When postgres client is notrunning properly", func(t *testing.T) {
		postgresClientMock.EXPECT().GetWallet(gomock.Any()).Do(func (reqDb database.GetWalletReqDb) {
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(database.GetWalletResDb{}, errSample)

		req := GetStatusReq{
			WalletId: "wallet-id-1",
		}
		res := &GetStatusRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.GetStatus(req, res)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, res.Status, "", "Status is not the same")

	})

}
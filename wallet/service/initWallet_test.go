package service

import (
	"testing"
	mock_database "wallet/database/mock_database"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"wallet/database"
	"errors"
)

func TestInitWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
    postgresClientMock := mock_database.NewMockPostgresClient(mockCtrl)
	errSample := errors.New("Error")

	
	t.Run("When postgres client is running properly", func(t *testing.T) {
		resDb := database.CreateWalletResDb{
			WalletId: "wallet-id-1",
		}
		postgresClientMock.EXPECT().CreateOrInsertWallet(gomock.Any()).Do(func (reqDb database.CreateWalletReqDb) {
			assert.Equal(t, reqDb.CustomerId, "customer-id-1", "Wallet is not the same")
			assert.Equal(t, reqDb.Status, "disabled", "Status is not the same")
		}).Return(resDb, nil)

		req := InitWalletReq{
			CustomerId: "customer-id-1",
		}
		res := &InitWalletRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.InitWallet(req, res)

		assert.Equal(t, err, nil, "Error should be nil")
		assert.NotNil(t, res.Token, "Token should not be nil")

	})

	t.Run("When postgres client is not running properly", func(t *testing.T) {
		postgresClientMock.EXPECT().CreateOrInsertWallet(gomock.Any()).Do(func (reqDb database.CreateWalletReqDb) {
			assert.Equal(t, reqDb.CustomerId, "customer-id-1", "Wallet is not the same")
			assert.Equal(t, reqDb.Status, "disabled", "Status is not the same")
		}).Return(database.CreateWalletResDb{}, errSample)

		req := InitWalletReq{
			CustomerId: "customer-id-1",
		}
		res := &InitWalletRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.InitWallet(req, res)

		assert.NotNil(t, err, "Error should not be nil")

	})

}
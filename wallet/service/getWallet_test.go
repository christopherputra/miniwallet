package service

import (
	"testing"
	mock_database "wallet/database/mock_database"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"wallet/database"
	"errors"
	funcs "wallet/functions"
)

func TestGetWallet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
    postgresClientMock := mock_database.NewMockPostgresClient(mockCtrl)
	errSample := errors.New("Error")

	resDb := database.ListTransactionsResDb{
		TransactionDbs: []database.TransactionDb{
			database.TransactionDb{
				TransactionId: "trans-1",
				Type: "deposit",
				WalletId: "wallet-1",
				Amount: 20000,
				ReferenceId: "reference-1",
				ExecutedAt: 1927412838121,
			},
			database.TransactionDb{
				TransactionId: "trans-2",
				Type: "withdraw",
				WalletId: "wallet-1",
				Amount: 100000,
				ReferenceId: "reference-2",
				ExecutedAt: 1927412838121,
			},
			database.TransactionDb{
				TransactionId: "trans-3",
				Type: "deposit",
				WalletId: "wallet-1",
				Amount: 180000,
				ReferenceId: "reference-3",
				ExecutedAt: 1927412838121,
			},
		},
	}
	postgresClientMock.EXPECT().ListTransactions(gomock.Any()).Do(func(reqDb database.ListTransactionsReqDb) {
		assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
	}).Return(resDb,nil)

	t.Run("When postgres client is not running properly", func(t *testing.T) {
		postgresClientMock.EXPECT().GetWallet(gomock.Any()).Do(func (reqDb database.GetWalletReqDb) {
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(database.GetWalletResDb{}, errSample)

		req := GetWalletReq{
			WalletId: "wallet-id-1",
		}
		res := &GetWalletRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.GetWallet(req, res)

		assert.NotNil(t, err, "Error should not be nil")
		
	})

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

		req := GetWalletReq{
			WalletId: "wallet-id-1",
		}
		res := &GetWalletRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.GetWallet(req, res)

		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, res.WalletData.OwnerId, "customer-id-1", "OwnerId is not the same")
		assert.Equal(t, res.WalletData.Balance, int64(100000), "Balance is not the same")
		assert.Equal(t, res.WalletData.EnabledAt, funcs.UnixToTimeStamp(18272491237), "Balance is not the same")

	})

}
package service

import (
	"testing"
	mock_database "wallet/database/mock_database"
	gomock "github.com/golang/mock/gomock"
	assert "github.com/stretchr/testify/assert"
	"wallet/database"
	"errors"
)

func TestListTransactions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
    postgresClientMock := mock_database.NewMockPostgresClient(mockCtrl)
	errSample := errors.New("Error")

	t.Run("When postgres client is not running properly", func(t *testing.T) {
		postgresClientMock.EXPECT().ListTransactions(gomock.Any()).Do(func(reqDb database.ListTransactionsReqDb) {
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(database.ListTransactionsResDb{},errSample)

		req := ListTransactionsReq{
			WalletId: "wallet-id-1",
		}
		res := &ListTransactionsRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.ListTransactions(req, res)

		assert.NotNil(t, err, "Error should not be nil")
	})
	
	t.Run("When postgres client is running properly", func(t *testing.T) {
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

		req := ListTransactionsReq{
			WalletId: "wallet-id-1",
		}
		res := &ListTransactionsRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.ListTransactions(req, res)

		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, len(res.Transactions), 3, "Length is not the same")
		assert.Equal(t, res.Transactions[0].TransactionId, "trans-1", "Length is not the same")
		assert.Equal(t, res.Transactions[1].TransactionId, "trans-2", "Length is not the same")
		assert.Equal(t, res.Transactions[2].TransactionId, "trans-3", "Length is not the same")
	})

}
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
func TestDepositTransaction(t *testing.T) {
	mockCtrl := gomock.NewController(t)
    postgresClientMock := mock_database.NewMockPostgresClient(mockCtrl)
	errSample := errors.New("Error")
	
	t.Run("When postgres client is running properly", func(t *testing.T) {
		resDb := database.InsertTransactionResDb{
			TransactionId: "trans-1",
			WalletId: "wallet-id-1",
			Type: "deposit",
			Amount: 20000,
			ReferenceId: "reference-1",
			ExecutedAt: 18272491237,
		}
		postgresClientMock.EXPECT().InsertTransaction(gomock.Any()).Do(func (reqDb database.InsertTransactionReqDb) {
			assert.Equal(t, reqDb.Type, "deposit", "Type is not Deposit")
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(resDb, nil)

		req := DepositTransactionReq{
			WalletId: "wallet-id-1",
		}
		res := &DepositTransactionRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.DepositTransaction(req, res)

		assert.Equal(t, err, nil, "Error should be nil")
		assert.Equal(t, res.Status, "success", "Status is not the same")
		assert.Equal(t, res.Amount, int64(20000), "Amount is not the same")
		assert.Equal(t, res.DepositedAt, funcs.UnixToTimeStamp(18272491237), "Amount is not the same")

	})

	t.Run("When postgres client is notrunning properly", func(t *testing.T) {
		postgresClientMock.EXPECT().InsertTransaction(gomock.Any()).Do(func (reqDb database.InsertTransactionReqDb) {
			assert.Equal(t, reqDb.Type, "deposit", "Type is not Deposit")
			assert.Equal(t, reqDb.WalletId, "wallet-id-1", "Wallet is not the same")
		}).Return(database.InsertTransactionResDb{}, errSample)

		req := DepositTransactionReq{
			WalletId: "wallet-id-1",
		}
		res := &DepositTransactionRes{}

		walletService := Service{PgServer: postgresClientMock}
		err := walletService.DepositTransaction(req, res)

		assert.NotNil(t, err, "Error should not be nil")
		assert.Equal(t, res.Status, "", "Status is not the same")

	})

}
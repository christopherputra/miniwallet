package functions

import (
	"testing"
	assert "github.com/stretchr/testify/assert"
)

func TestTokenValidation(t *testing.T) {
	walletId := "wallet-1"
	tokenString, err := GenerateToken(walletId)
	assert.Equal(t, err, nil, "Should be nil")
	valid, walletIdParsed, err := VerifyToken(tokenString)
	assert.Equal(t, err, nil, "Should be nil")
	assert.Equal(t, walletIdParsed, walletId, "Wallet Id is different")
	assert.Equal(t, valid, true, "Token is invalid")

	walletId = "wallet-2"
	tokenString, err = GenerateToken(walletId)
	assert.Equal(t, err, nil, "Should be nil")
	valid, walletIdParsed, err = VerifyToken(tokenString)
	assert.Equal(t, err, nil, "Should be nil")
	assert.Equal(t, walletIdParsed, walletId, "Wallet Id is different")
	assert.Equal(t, valid, true, "Token should be valid")

	//IF TOKEN MUTATED THEN SHOULD BE INVALID
	walletId = "wallet-3"
	tokenString, err = GenerateToken(walletId)
	assert.Equal(t, err, nil, "Should be nil")
	//MUTATE
	tokenString = tokenString+"af345"
	valid, walletIdParsed, err = VerifyToken(tokenString)
	assert.NotNil(t, err, "Should be not nil")
	assert.Equal(t, valid, false, "Token should be false")
}
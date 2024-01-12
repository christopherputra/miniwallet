package database


type CreateWalletReqDb struct {
	WalletId string
	CustomerId string
	Status string
}
type CreateWalletResDb struct {
	WalletId string `db:"id"`
}

func (pg *PostgresServer) CreateOrInsertWallet (req CreateWalletReqDb) (CreateWalletResDb, error) {
	res := CreateWalletResDb {}
	rows, err := pg.Db.Query(`INSERT INTO wallets (customer_id, id, status)
	VALUES ($1, $2, $3)
	ON CONFLICT (customer_id) DO UPDATE SET id=wallets.id RETURNING id`, req.CustomerId, req.WalletId, req.Status)
	if err != nil {
		return CreateWalletResDb{}, err
	}
	for rows.Next() {
		err := rows.Scan(&res.WalletId)
		if err != nil {
			return CreateWalletResDb{}, err
		} 
	}
	return res, nil
}

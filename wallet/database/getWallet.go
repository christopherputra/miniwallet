package database


type GetWalletReqDb struct {
	WalletId string
}
type GetWalletResDb struct {
	WalletId string `db:"id"`
	CustomerId string `db:"customer_id"`
	Status string `db:"status"`
	EnabledAt int64 `db:"enabled_at"`
}

func (pg *PostgresServer) GetWallet (req GetWalletReqDb) (GetWalletResDb, error) {
	res := GetWalletResDb {}
	rows, err := pg.Db.Query(`SELECT id, customer_id, status, enabled_at FROM wallets WHERE id = $1`, req.WalletId)
	if err != nil {
		return GetWalletResDb{}, err
	}
	for rows.Next() {
		err := rows.Scan(&res.WalletId, &res.CustomerId, &res.Status, &res.EnabledAt)
		if err != nil {
			return GetWalletResDb{}, err
		} 
	}
	return res, nil
}

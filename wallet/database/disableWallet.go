package database


type DisableWalletReqDb struct {
	WalletId string
	Status string
	DisabledAt int64
}
type DisableWalletResDb struct {
	WalletId string `db:"id"`
	CustomerId string `db:"customer_id"`
	Status string `db:"status"`
	DisabledAt int64 `db:"enabled_at"`
}

func (pg *PostgresServer) DisableWallet (req DisableWalletReqDb) (DisableWalletResDb, error) {
	res := DisableWalletResDb {}
	rows, err := pg.Db.Query(`UPDATE wallets SET 
	status = $1, disabled_at = $2 WHERE id = $3 RETURNING id, customer_id, status, disabled_at`, req.Status, req.DisabledAt, req.WalletId)
	if err != nil {
		return DisableWalletResDb{}, err
	}
	for rows.Next() {
		err := rows.Scan(&res.WalletId, &res.CustomerId, &res.Status, &res.DisabledAt)
		if err != nil {
			return DisableWalletResDb{}, err
		} 
	}
	return res, nil
}

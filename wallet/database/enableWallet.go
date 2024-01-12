package database


type EnableWalletReqDb struct {
	WalletId string
	Status string
	EnabledAt int64
}
type EnableWalletResDb struct {
	WalletId string `db:"id"`
	CustomerId string `db:"customer_id"`
	Status string `db:"status"`
	EnabledAt int64 `db:"enabled_at"`
}

func (pg *PostgresServer) EnableWallet (req EnableWalletReqDb) (EnableWalletResDb, error) {
	res := EnableWalletResDb {}
	rows, err := pg.Db.Query(`UPDATE wallets SET 
	status = $1, enabled_at = $2 WHERE id = $3 RETURNING id, customer_id, status, enabled_at`, req.Status, req.EnabledAt, req.WalletId)
	if err != nil {
		return EnableWalletResDb{}, err
	}
	for rows.Next() {
		err := rows.Scan(&res.WalletId, &res.CustomerId, &res.Status, &res.EnabledAt)
		if err != nil {
			return EnableWalletResDb{}, err
		} 
	}
	return res, nil
}

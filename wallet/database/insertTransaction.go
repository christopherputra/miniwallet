package database


type InsertTransactionReqDb struct {
	TransactionId string
	Type string
	WalletId string
	Amount int64
	ReferenceId string
	ExecutedAt int64
}
type InsertTransactionResDb struct {
	TransactionId string
	Type string
	WalletId string
	Amount int64
	ReferenceId string
	ExecutedAt int64
}

func (pg *PostgresServer) InsertTransaction (req InsertTransactionReqDb) (InsertTransactionResDb, error) {
	res := InsertTransactionResDb {}
	rows, err := pg.Db.Query(`INSERT INTO transactions (id, type, wallet_id, amount, reference_id, executed_at)
	VALUES ($1,$2,$3,$4,$5,$6) RETURNING id, type, wallet_id, amount, reference_id, executed_at`, req.TransactionId, req.Type, req.WalletId, req.Amount, req.ReferenceId, req.ExecutedAt)
	if err != nil {
		return InsertTransactionResDb{}, err
	}
	for rows.Next() {
		err := rows.Scan(&res.TransactionId, &res.Type, &res.WalletId, &res.Amount, &res.ReferenceId, &res.ExecutedAt)
		if err != nil {
			return InsertTransactionResDb{}, err
		} 
	}
	return res, nil
}

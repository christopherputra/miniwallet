package database


type ListTransactionsReqDb struct {
	WalletId string
}

type TransactionDb struct {
	TransactionId string `db:"id"`
	Type string `db:"type"`
	WalletId string `db:"wallet_id"`
	Amount int64 `db:"amount"`
	ReferenceId string `db:"reference_id"`
	ExecutedAt int64 `db:"executed_at"`
}

type ListTransactionsResDb struct {
	TransactionDbs []TransactionDb
}

func (pg *PostgresServer) ListTransactions (req ListTransactionsReqDb) (ListTransactionsResDb, error) {
	res := ListTransactionsResDb {}
	rows, err := pg.Db.Query(`SELECT id, type, wallet_id, amount, reference_id, executed_at FROM transactions WHERE wallet_id = $1`, req.WalletId)
	if err != nil {
		return ListTransactionsResDb{}, err
	}
	transactions := make([]TransactionDb, 0)
	for rows.Next() {
		rowResult := TransactionDb{}
		err := rows.Scan(&rowResult.TransactionId, &rowResult.Type, &rowResult.WalletId, &rowResult.Amount, &rowResult.ReferenceId, &rowResult.ExecutedAt)
		if err != nil {
			return ListTransactionsResDb{}, err
		}
		transactions = append(transactions, rowResult)
	}
	res.TransactionDbs = transactions
	return res, nil
}

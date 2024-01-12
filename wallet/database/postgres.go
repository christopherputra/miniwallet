package database

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresClient interface {
    CreateOrInsertWallet (req CreateWalletReqDb) (CreateWalletResDb, error)
    EnableWallet (req EnableWalletReqDb) (EnableWalletResDb, error)
    DisableWallet (req DisableWalletReqDb) (DisableWalletResDb, error)
    GetWallet (req GetWalletReqDb) (GetWalletResDb, error)
    InsertTransaction (req InsertTransactionReqDb) (InsertTransactionResDb, error)
    ListTransactions (req ListTransactionsReqDb) (ListTransactionsResDb, error)
}
type PostgresServer struct {
    Db *sql.DB
}
func NewPostgresClient(host string, port int, user string, password string, dbname string) PostgresClient {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    
	dbserver, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
    return &PostgresServer{
        Db: dbserver,
    }
}
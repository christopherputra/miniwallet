# Mini Wallet App Case in Golang
### About
mini wallet app is an app that works as a virtual wallet. you can initiate your wallet, then enable the wallet to use its features. Features of the wallet are: 
1. View Wallet with its final amount balance
2. Add transactions (Deposit/Withdraw)
3. View history transactions

### Assumption and Disclaimer
1.**JWT TOKEN SIGNING** I don't use standard Tokenization, instead I use JWT signing system to improve the flow and reduce processing time. Using the standard token verification that need to be stored in the database is not effective due to the need of database access in order to get the token for every matchup (eg. need to access the database everytime a request is made, best is using a caching db like Redis but still not effective). On the other hand, by using JWT token signing, Token are constructed of payload (any data  we want) that were encoded using a secret private key. So everytime a request is made the token only need to be decoded back using our private key, so no database access is needed. if the token is mutated or added by unauthorized user, the token will not be able to be decoded using our secret private key, hence considered Unauthorized. Moreover, JWT supports any payload added to the token, which in this case is the Wallet Id added to the payload. This is super effective and also secure compared to the standard token system.
### Brief Architecture - Service Based API
[![ARCHITECTURE](https://i.imgur.com/lmNhenG.png "ARCHITECTURE")](https://i.imgur.com/lmNhenG.png "ARCHITECTURE")
### Stacks
- Golang
- JWT Token Signing https://pkg.go.dev/github.com/golang-jwt/jwt/v5
- Go Gin Routing https://github.com/gin-gonic/gin
- Postgresql

| [![go](https://i.imgur.com/miVUk6U.png "go")](https://i.imgur.com/miVUk6U.png "go")  | [![go gin](https://i.imgur.com/8OTwAo4.png "go gin")](https://i.imgur.com/8OTwAo4.png "go gin")  |  [![jwt](https://i.imgur.com/2GujZmD.png "jwt")](https://i.imgur.com/2GujZmD.png "jwt") |  [![postgres](https://i.imgur.com/dLxfiGU.png "postgres")](https://i.imgur.com/dLxfiGU.png "postgres") |
| ------------ | ------------ | ------------ | ------------ |
|   |   |   |   |    |
### Code Structure
```
wallet
│
└───database
│   │  accessorName.go
│   │  accessorName2.go
│   │   ...
│   │
└───functions
│   │   function1.go
│   │   function2.go
│   │   ...
│   
└───middlewares
│   │   middleware1.go
│   │   middleware2.go
│   │  ...
└───service
│   │   service.go
│   │   endpoints1.go
│   │   endpoints2.go
│   │   ...
│   │
│   └───handler
│       │   handler.go
│   
│   config.json
│   go.mod
│   go.sum
│   main.go
```

## Getting Started
#### Golang Install
Download Go https://go.dev/doc/install

#### PostgreSQL Install
Please install postgresql on your local following these steps https://www.postgresql.org/download/macosx/

## First Setup
#### PostgreSQL Schema
make `wallet` database
run `schema.sql` on your postgres server to create all tables
```sql
DROP TABLE IF EXISTS wallets;
CREATE TABLE wallets (
	id VARCHAR ( 36 ) UNIQUE NOT NULL,
	customer_id VARCHAR ( 36 ) UNIQUE NOT NULL,
	status VARCHAR ( 10 ) NOT NULL,
	enabled_at BIGINT,
	disabled_at BIGINT
);

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
	id VARCHAR ( 36 ) UNIQUE NOT NULL,
	type VARCHAR ( 10 ) NOT NULL,
	wallet_id VARCHAR ( 36 ) NOT NULL,
	amount INT NOT NULL,
	reference_id VARCHAR ( 36 ) UNIQUE NOT NULL,
	executed_at BIGINT NOT NULL
);
```

#### Setup config.json File
Change env fields based on your local env including postgres host, port, username, and password.
```json
{
    "postgres": {
        "host": "localhost",
        "port": 5432,
        "user": "<username>",
        "password": "<password>",
        "dbname": "wallet"
    },
    "host": "localhost",
    "port": "8080"
}
```
#### Go Get Packages
Please `go get <packages not yet installed>`

## You're Good To GO - Run Service
on your console inside `/wallet` dir
run go server using this command
`go run main.go`

## Unit Tests
All of the endpoints are already tested in unit granularity.
if you wish to run the test, go to the directive folder with the tests files and run `go test`



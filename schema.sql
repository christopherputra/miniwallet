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


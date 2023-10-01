package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"time"
)

type TransactionsQ interface {
	New() TransactionsQ
	Insert(data *Transaction) error
	Select() ([]Transaction, error)
	Page(pageParams pgdb.OffsetPageParams) TransactionsQ
	Update(data *Transaction) error
	FilterByRecipient(address string) TransactionsQ
	OrderBy(column, order string) TransactionsQ
}

type Transaction struct {
	Hash        string    `db:"hash" structs:"hash"`
	Recipient   string    `db:"recipient" structs:"recipient"`
	Sender      string    `db:"sender" structs:"sender"`
	ValueTo     string    `db:"value_to" structs:"value_to"`
	Currency    string    `db:"currency" structs:"currency"`
	TimestampTo time.Time `db:"timestamp_to" structs:"timestamp_to"`
}

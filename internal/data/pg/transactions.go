package pg

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const transactionTableName = "transactions"

const (
	ID             = "id"
	RecipientField = "recipient"
	Hash           = "hash"
	TimestampField = "timestamp_to"
)

func NewTransactionsQ(db *pgdb.DB) data.TransactionsQ {
	return &TransactionsQ{
		db:  db,
		sql: sq.Select("b.*").From(fmt.Sprintf("%s as b", transactionTableName)),
		upd: sq.Update(transactionTableName),
	}
}

type TransactionsQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *TransactionsQ) New() data.TransactionsQ {
	return NewTransactionsQ(q.db.Clone())
}

func (q *TransactionsQ) Update(data *data.Transaction) error {
	clauses := structs.Map(data)
	if err := q.db.Exec(q.upd.SetMap(clauses)); err != nil {
		return errors.Wrap(err, "failed to update data")
	}

	return nil
}

func (q *TransactionsQ) Select() ([]data.Transaction, error) {
	var result []data.Transaction
	err := q.db.Select(&result, q.sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select txs")
	}

	return result, nil
}

func (q *TransactionsQ) Insert(value *data.Transaction) error {
	clauses := structs.Map(value)

	if err := q.db.Exec(sq.Insert(transactionTableName).SetMap(clauses)); err != nil {
		return errors.Wrap(err, "failed to insert tx")
	}

	return nil
}

func (q *TransactionsQ) FilterByRecipient(address string) data.TransactionsQ {
	q.sql = q.sql.Where(sq.Eq{RecipientField: address})
	q.upd = q.upd.Where(sq.Eq{RecipientField: address})

	return q
}

func (q *TransactionsQ) OrderBy(column, order string) data.TransactionsQ {
	q.sql = q.sql.OrderBy(fmt.Sprintf("%s %s", column, order))
	q.upd = q.upd.OrderBy(fmt.Sprintf("%s %s", column, order))

	return q
}

func (q *TransactionsQ) Page(pageParams pgdb.OffsetPageParams) data.TransactionsQ {
	q.sql = pageParams.ApplyTo(q.sql, ID)

	return q
}

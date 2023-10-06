package pg

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const datasetTableName = "dataset"

const (
	ID             = "id"
	TokenNameField = "tokenName"
)

func NewDatasetQ(db *pgdb.DB) data.DatasetQ {
	return &DatasetQ{
		db:  db,
		sql: sq.Select("b.*").From(fmt.Sprintf("%s as b", datasetTableName)),
		upd: sq.Update(datasetTableName),
	}
}

type DatasetQ struct {
	db  *pgdb.DB
	sql sq.SelectBuilder
	upd sq.UpdateBuilder
}

func (q *DatasetQ) New() data.DatasetQ {
	return NewDatasetQ(q.db.Clone())
}

func (q *DatasetQ) Update(data *data.Dataset) error {
	clauses := structs.Map(data)
	if err := q.db.Exec(q.upd.SetMap(clauses)); err != nil {
		return errors.Wrap(err, "failed to update data")
	}

	return nil
}

func (q *DatasetQ) Select() ([]data.Dataset, error) {
	var result []data.Dataset
	err := q.db.Select(&result, q.sql)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select txs")
	}

	return result, nil
}

func (q *DatasetQ) Insert(value *data.Dataset) error {
	clauses := structs.Map(value)

	if err := q.db.Exec(sq.Insert(datasetTableName).SetMap(clauses)); err != nil {
		return errors.Wrap(err, "failed to insert tx")
	}

	return nil
}

func (q *DatasetQ) FilterByTokenName(tokenName string) data.DatasetQ {
	q.sql = q.sql.Where(sq.Eq{TokenNameField: tokenName})
	q.upd = q.upd.Where(sq.Eq{TokenNameField: tokenName})

	return q
}

func (q *DatasetQ) OrderBy(column, order string) data.DatasetQ {
	q.sql = q.sql.OrderBy(fmt.Sprintf("%s %s", column, order))
	q.upd = q.upd.OrderBy(fmt.Sprintf("%s %s", column, order))

	return q
}

func (q *DatasetQ) Page(pageParams pgdb.OffsetPageParams) data.DatasetQ {
	q.sql = pageParams.ApplyTo(q.sql, ID)

	return q
}

package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

type DatasetQ interface {
	New() DatasetQ
	Insert(data *Dataset) error
	Select() ([]Dataset, error)
	Page(pageParams pgdb.OffsetPageParams) DatasetQ
	Update(data *Dataset) error
	FilterByTokenName(address string) DatasetQ
	OrderBy(column, order string) DatasetQ
}

type Dataset struct {
	TokenName            string `db:"token_name" structs:"token_name"`
	TxsNumber            string `db:"txs_number" structs:"txs_number"`
	AverageTxsCount      string `db:"average_txs_count" structs:"average_txs_count"`
	AverageGoogleSites   string `db:"average_google_sites" structs:"average_google_sites"`
	IsThereGithub        string `db:"is_there_github" structs:"is_there_github"`
	TokenDescription     string `db:"token_description" structs:"token_description"`
	NumberOfUserGroup    string `db:"number_of_user_group" structs:"number_of_user_group"`
	PercentTokenHandlers string `db:"percent_token_handlers" structs:"percent_token_handlers"`
}

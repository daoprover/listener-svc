package data

type MasterQ interface {
	New() MasterQ
	TransactionsQ() DatasetQ
	Transaction(func(data interface{}) error, interface{}) error
}

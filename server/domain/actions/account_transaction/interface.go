package account_transaction

type IAccountTransaction interface {
	UpdateValue(id *int64, value *float64) (err error)
}

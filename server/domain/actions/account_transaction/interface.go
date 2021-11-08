package account_transaction

type IAccountTransaction interface {
	Deposit(id *int64, value *float64) (err error)
}

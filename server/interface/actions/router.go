package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/interface/actions/account"
	"github.com/hiagomf/bank-api/server/interface/actions/account_detail"
	"github.com/hiagomf/bank-api/server/interface/actions/account_transaction"
	"github.com/hiagomf/bank-api/server/interface/actions/payment_slip"
)

func Router(r *gin.RouterGroup) {
	account.Router(r.Group("accounts"))
	account.RouterID(r.Group("account"))

	account_detail.Router(r.Group("account_detail"))

	account_transaction.Router(r.Group("account_transactions"))

	payment_slip.Router(r.Group("payment_slips"))
}

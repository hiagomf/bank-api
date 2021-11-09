package records

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/interface/records/account_owner"
	"github.com/hiagomf/bank-api/server/interface/records/account_owner_address"
	"github.com/hiagomf/bank-api/server/interface/records/bank"
	"github.com/hiagomf/bank-api/server/interface/records/payment_slip"
)

func Router(r *gin.RouterGroup) {
	bank.Router(r.Group("banks"))
	bank.RouterID(r.Group("bank"))

	account_owner.Router(r.Group("accounts_owners"))
	account_owner.RouterID(r.Group("account_owner"))

	account_owner_address.Router(r.Group("account_owner_addresses"))
	account_owner_address.RouterID(r.Group("account_owner_address"))

	payment_slip.Router(r.Group("payment_slips"))
}

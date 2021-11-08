package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/interface/actions/account"
	"github.com/hiagomf/bank-api/server/interface/actions/account_detail"
)

func Router(r *gin.RouterGroup) {
	account.Router(r.Group("accounts"))
	account.RouterID(r.Group("account"))

	account_detail.Router(r.Group("account_detail"))
}

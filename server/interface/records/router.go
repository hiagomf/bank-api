package records

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/interface/records/bank"
)

func Router(r *gin.RouterGroup) {
	bank.Router(r.Group("banks"))
	bank.RouterID(r.Group("bank"))
}

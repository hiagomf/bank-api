package bank

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/interface/records/agency"
)

func Router(r *gin.RouterGroup) {
	r.GET("", selectPaginated)
}

func RouterID(r *gin.RouterGroup) {
	r.GET(":bank_id/agencies", agency.SelectPaginatedByBank)
}

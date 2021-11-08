package account_transaction

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("deposit", deposit)
}

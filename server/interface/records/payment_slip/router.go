package payment_slip

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("", selectPaginated)
}

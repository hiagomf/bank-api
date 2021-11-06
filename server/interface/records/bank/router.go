package bank

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", selectPaginated)
}

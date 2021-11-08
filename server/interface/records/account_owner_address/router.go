package account_owner_address

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", selectPaginated)
	r.POST("", insert)
}

func RouterID(r *gin.RouterGroup) {
	r.GET(":address_id", selectOne)
	r.PUT(":address_id", update)
	r.DELETE(":address_id", disable)
}

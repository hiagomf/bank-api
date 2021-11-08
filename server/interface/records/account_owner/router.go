package account_owner

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.GET("", selectPaginated)
	r.POST("", insert)
}

func RouterID(r *gin.RouterGroup) {
	r.GET(":owner_id", selectOne)
	r.PUT(":owner_id", update)
	r.DELETE(":owner_id", disable)
}

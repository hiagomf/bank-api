package account

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("", openAccount)
}

func RouterID(r *gin.RouterGroup) {
	r.DELETE(":account_id", closeAccount)
}

package account_detail

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/actions/account_detail"
	"github.com/hiagomf/bank-api/server/oops"
)

func checkDetail(c *gin.Context) {
	var req account_detail.Request

	err := c.ShouldBindJSON(&req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := account_detail.CheckDetails(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

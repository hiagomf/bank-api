package account

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/actions/account"
	"github.com/hiagomf/bank-api/server/oops"
)

func openAccount(c *gin.Context) {
	var req account.Request

	err := c.ShouldBindJSON(&req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := account.OpenAccount(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

func closeAccount(c *gin.Context) {
	paramID := c.Param("account_id")
	id, err := strconv.ParseInt(paramID, 10, 0)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := account.CloseAccount(c, &id); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

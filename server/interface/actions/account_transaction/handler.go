package account_transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/actions/account_transaction"
	"github.com/hiagomf/bank-api/server/oops"
)

func deposit(c *gin.Context) {
	var req account_transaction.DepositRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := account_transaction.Deposit(c, &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

func transfer(c *gin.Context) {
	var req account_transaction.TransferRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := account_transaction.Transfer(c, &req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

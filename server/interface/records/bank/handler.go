package bank

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/records/bank"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

func selectPaginated(c *gin.Context) {
	p, err := utils.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	res, err := bank.SelectPaginated(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

package agency

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/records/agency"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

func SelectPaginatedByBank(c *gin.Context) {
	param := c.Param("bank_id")

	bankID, err := strconv.ParseInt(param, 10, 0)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	p, err := utils.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	res, err := agency.SelectPaginatedByBank(c, &bankID, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

package account_owner_address

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/records/account_owner_address"
	"github.com/hiagomf/bank-api/server/oops"
	"github.com/hiagomf/bank-api/server/utils"
)

func insert(c *gin.Context) {
	var req account_owner_address.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := account_owner_address.Insert(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

func update(c *gin.Context) {
	var req account_owner_address.Request

	paramID := c.Param("address_id")
	id, err := strconv.ParseInt(paramID, 10, 0)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := account_owner_address.Update(c, &req, &id); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

func selectPaginated(c *gin.Context) {
	p, err := utils.ParseParams(c)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	res, err := account_owner_address.SelectPaginated(c, &p)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

func selectOne(c *gin.Context) {
	paramID := c.Param("address_id")
	id, err := strconv.ParseInt(paramID, 10, 0)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	res, err := account_owner_address.SelectOne(c, &id)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(200, res)
}

func disable(c *gin.Context) {
	paramID := c.Param("address_id")
	id, err := strconv.ParseInt(paramID, 10, 0)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	if err := account_owner_address.Disable(c, &id); err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(204, nil)
}

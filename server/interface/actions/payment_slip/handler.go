package payment_slip

import (
	"github.com/gin-gonic/gin"
	"github.com/hiagomf/bank-api/server/application/actions/payment_slip"
	"github.com/hiagomf/bank-api/server/oops"
)

func generate(c *gin.Context) {
	var req payment_slip.Request

	err := c.ShouldBindJSON(&req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	id, err := payment_slip.GeneratePaymentSlip(c, &req)
	if err != nil {
		oops.DefinirErro(err, c)
		return
	}

	c.JSON(201, id)
}

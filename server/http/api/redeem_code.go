package api

import (
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/http/data"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/log"

	"github.com/gin-gonic/gin"
)

// GenerateRedeemCodes godoc
// @Summary Generate redeem codes
// @Description Generate redeem codes
// @Tags RedeemCodes
// @Accept json
// @Produce json
// @Param amount query int true "Amount for each redeemm code"
// @Param count query int true "Number of redeem codes to generate"
// @Success 200 {object} data.BaseResponse{data=data.RedeemCodeGenResult}
// @Failure 400 {object} data.BaseResponse
// @Router /payment-ms/v1/merchant/redeem-codes/generate [post]
func GenerateRedeemCodes(c *gin.Context) {
	var req data.RedeemCodeGenRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Logger.Errorf("GenerateRedeemCodes bind error:", err)
		c.JSON(400, data.BaseResponse{ErrMsg: err.Error()})
		return
	}
	ret := &data.RedeemCodeGenResult{GenCount: 0}
	c.JSON(200, data.BaseResponse{Data: ret})
}

// QueryRedeemCodes godoc
// @Summary Query redeem codes
// @Description Query redeem codes
// @Tags RedeemCodes
// @Accept json
// @Produce json
// @Param query query data.RedeemCodeQuery false "Redeem code to search for"
// Success 200 {object} data.BaseResponse{data=[]data.RedeemCodeVO}
// @Failure 400 {object} data.BaseResponse
// @Router /payment-ms/v1/merchant/redeem-codes [get]
func QueryRedeemCodes(c *gin.Context) {
	var query data.RedeemCodeQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Logger.Errorf("QueryRedeemCodes bind error:", err)
		c.JSON(400, data.BaseResponse{ErrMsg: err.Error()})
		return
	}
	ret := make([]data.RedeemCodeVO, 0)
	c.JSON(200, data.BaseResponse{Data: ret})
}

package api

import (
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/http/data"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/log"
	"github.com/gin-gonic/gin"
)

// GetUserPayAccountInfo godoc
// @Summary Get user pay account info
// @Description Get user pay account info
// @Tags PayAccount
// @Accept json
// @Produce json
// @Success 200 {object} data.BaseResponse{data=data.UserPayAccount}
// @Failure 400 {object} data.BaseResponse
// @Router /payment-ms/v1/customer/pay-accounts/self [get]
func GetUserPayAccountInfo(c *gin.Context) {
	var userId int
	if userIdInterface, exists := c.Get("userID"); exists {
		userId = userIdInterface.(int)
	} else {
		c.JSON(400, gin.H{"error": "User ID not found in context"})
		return
	}

	// Simulate fetching user pay account info from database
	payAccount := &data.UserPayAccount{
		UserId:    userId,
		AccountNo: "ACC123456",
		Balance:   1000,
		CreatedAt: 1622547800,
		UpdatedAt: 1625149800,
	}

	c.JSON(200, data.BaseResponse{Data: payAccount})
}

// TopUpUserPayAccount godoc
// @Summary Top up user pay account
// @Description Top up user pay account using a redeem code
// @Tags PayAccount
// @Accept json
// @Produce json
// @Param topup body data.UserPayAccountTopUpRequest true "Top up request"
// @Success 200 {object} data.BaseResponse{data=data.UserPayAccountTopUpResult}
// @Failure 400 {object} data.BaseResponse
// @Router /payment-ms/v1/customer/pay-accounts/self/top-ups [post]
func TopUpUserPayAccount(c *gin.Context) {
	var userId int
	if userIdInterface, exists := c.Get("userID"); exists {
		userId = userIdInterface.(int)
	} else {
		c.JSON(400, gin.H{"error": "User ID not found in context"})
		return
	}
	log.Logger.Infof("TopUpUserPayAccount called for user ID: %d", userId)
	var req data.UserPayAccountTopUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, data.BaseResponse{ErrMsg: err.Error()})
		return
	}

	// Simulate processing the top-up request
	topUpAmount := 100 // This would be determined by the redeem code in a real implementation
	currentBalance := 1100

	result := &data.UserPayAccountTopUpResult{
		TopUpAmount:    topUpAmount,
		CurrentBalance: currentBalance,
	}

	c.JSON(200, data.BaseResponse{Data: result})
}

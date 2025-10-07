package router

import (
	"github.com/gin-gonic/gin"

	_ "github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/docs"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-payment-mservice/server/http/api"
	"github.com/NUS-ISS-Agile-Team/ceramicraft-user-mservice/common/middleware"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

const (
	serviceURIPrefix = "/payment-ms/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	basicGroup := r.Group(serviceURIPrefix)
	{
		basicGroup.GET("/swagger/*any", gs.WrapHandler(
			swaggerFiles.Handler,
			gs.URL("/payment-ms/v1/swagger/doc.json"),
		))
		basicGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	v1Authed := basicGroup.Group("")
	{
		v1Authed.Use(middleware.AuthMiddleware())
		v1Authed.GET("/merchant/redeem-codes", api.QueryRedeemCodes)
		v1Authed.POST("/merchant/redeem-codes/generate", api.GenerateRedeemCodes)
		v1Authed.POST("/customer/pay-accounts/self/top-ups", api.TopUpUserPayAccount)
		v1Authed.GET("/customer/pay-accounts/self", api.GetUserPayAccountInfo)
	}
	return r
}

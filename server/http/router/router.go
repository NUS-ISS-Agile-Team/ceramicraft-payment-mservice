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

	v1UnAuthed := r.Group(serviceURIPrefix + "/:client")
	{
		v1UnAuthed.Use(validateClient())
		v1UnAuthed.POST("/items", api.CreateItem)
		v1UnAuthed.GET("/items/:item_id", api.GetItems)
	}
	v1Authed := r.Group(serviceURIPrefix + "/:client")
	{
		v1Authed.Use(validateClient(), middleware.AuthMiddleware())
		//todo: add authed api
	}
	return r
}

func validateClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := c.Param("client")
		if client != "merchant" && client != "customer" {
			c.JSON(400, gin.H{"error": "Invalid client type"})
			c.Abort()
			return
		}
		c.Next()
	}
}

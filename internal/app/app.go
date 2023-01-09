package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleErrors(c *gin.Context) {
	c.Next()
	if errToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last(); errToPrint != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": errToPrint.Error(),
		})
	}
}

// Register API endpoints to the router
func SetupRoutes(r *gin.Engine, ctx context.Context) {
	c := NewBotController(ctx)
	r.Use(handleErrors)
	r.GET("/api/v1/linebot/messages", c.get)
	r.POST("/api/v1/linebot/messages", c.send)
	r.POST("/api/v1/linebot/webhook", c.webhook)
}

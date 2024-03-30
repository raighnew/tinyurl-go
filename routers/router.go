package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raighneweng/tinyurl-go/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/:urlHash", api.GetFullUrl)

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/generate", api.GenerateShortURL)

	return r
}

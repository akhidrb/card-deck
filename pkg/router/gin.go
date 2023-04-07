package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Init(dbConn *gorm.DB) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(gin.DefaultWriter), gin.Recovery())
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	apiV1 := engine.Group("/api/v1")
	InitDeckDependencies(apiV1, dbConn)
	return engine
}

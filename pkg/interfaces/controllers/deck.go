package controllers

import "github.com/gin-gonic/gin"

type IDeck interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
}

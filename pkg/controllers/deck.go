package controllers

import (
	"github.com/akhidrb/toggl-cards/pkg/dtos"
	controllersI "github.com/akhidrb/toggl-cards/pkg/interfaces/controllers"
	servicesI "github.com/akhidrb/toggl-cards/pkg/interfaces/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Deck struct {
	service servicesI.IDeck
}

func NewDeck(service servicesI.IDeck) controllersI.IDeck {
	return Deck{service: service}
}

func (ctrl Deck) Create(c *gin.Context) {
	request := dtos.CreateDeckRequest{}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	response, err := ctrl.service.Create(request)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"message": "server error",
			})
		return
	}
	c.JSON(http.StatusCreated, response)
}

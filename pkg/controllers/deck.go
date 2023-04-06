package controllers

import (
	controllersI "github.com/akhidrb/toggl-cards/pkg/interfaces/controllers"
	servicesI "github.com/akhidrb/toggl-cards/pkg/interfaces/services"
	"github.com/gin-gonic/gin"
)

type Deck struct {
	service servicesI.IDeck
}

func NewDeck(service servicesI.IDeck) controllersI.IDeck {
	return Deck{service: service}
}

func (ctrl Deck) Create(c *gin.Context) {
	return
}

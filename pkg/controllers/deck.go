package controllers

import (
	"github.com/akhidrb/toggl-cards/pkg/config"
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
		config.HandleResponseError(c, config.NewBadRequestError(err))
		return
	}
	response, err := ctrl.service.Create(request)
	if err != nil {
		config.HandleResponseError(c, err)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (ctrl Deck) GetByID(c *gin.Context) {
	request := dtos.OpenDeckRequest{}
	if err := c.ShouldBindUri(&request); err != nil {
		config.HandleResponseError(c, config.NewBadRequestError(err))
		return
	}
	response, err := ctrl.service.GetByID(request)
	if err != nil {
		config.HandleResponseError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ctrl Deck) DrawCards(c *gin.Context) {
	uri := dtos.DrawCardsRequestURI{}
	if err := c.ShouldBindUri(&uri); err != nil {
		config.HandleResponseError(c, config.NewBadRequestError(err))
		return
	}
	params := dtos.DrawCardsRequestParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		config.HandleResponseError(c, config.NewBadRequestError(err))
		return
	}
	request := dtos.DrawCardsRequest{
		DrawCardsRequestURI:    uri,
		DrawCardsRequestParams: params,
	}
	response, err := ctrl.service.DrawCards(request)
	if err != nil {
		config.HandleResponseError(c, err)
		return
	}
	c.JSON(http.StatusOK, response)
}

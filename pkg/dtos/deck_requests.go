package dtos

import (
	"github.com/akhidrb/toggl-cards/pkg/models"
)

type CreateDeckRequest struct {
	Shuffle   bool    `form:"shuffle"`
	Cards     *string `form:"cards"`
	CardsList []string
}

func (dto *CreateDeckRequest) ToModel() models.Deck {
	return models.Deck{
		Shuffled: dto.Shuffle,
		Cards:    dto.CardsList,
	}
}

type OpenDeckRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type DrawCardsRequest struct {
	DrawCardsRequestURI
	DrawCardsRequestParams
}

type DrawCardsRequestURI struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type DrawCardsRequestParams struct {
	Count int `form:"count"`
}

func (dto *DrawCardsRequest) UpdateModel(deck *models.Deck) {
	deck.Cards = deck.Cards[dto.Count:]
}

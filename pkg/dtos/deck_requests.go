package dtos

import (
	"errors"
	"fmt"
	"github.com/akhidrb/toggl-cards/pkg/config"
	"github.com/akhidrb/toggl-cards/pkg/models"
	"strings"
)

type CreateDeckRequest struct {
	Shuffle   bool    `form:"shuffle"`
	Cards     *string `form:"cards"`
	CardsList []string
}

func (dto *CreateDeckRequest) ValidateCards() (err error) {
	for _, card := range dto.CardsList {
		cardKeys := strings.Split(card, "")
		rankKey := cardKeys[0]
		suitKey := cardKeys[1]
		err = config.NewValidationError(
			errors.New(
				fmt.Sprintf(
					"card %s is not part of the 52 card deck", card,
				),
			),
		)
		if _, ok := cardRanksMap[rankKey]; !ok {
			return err
		}
		if _, ok := cardSuitsMap[suitKey]; !ok {
			return err
		}
	}
	return nil
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

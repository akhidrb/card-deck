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

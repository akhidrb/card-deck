package dtos

import (
	"github.com/akhidrb/toggl-cards/pkg/models"
)

type CreateDeckRequest struct {
	Shuffle bool     `json:"shuffle"`
	Cards   []string `json:"cards"`
}

func (dto *CreateDeckRequest) ToModel() models.Deck {
	return models.Deck{
		Shuffled: dto.Shuffle,
		Cards:    dto.Cards,
	}
}

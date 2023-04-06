package dtos

import "github.com/akhidrb/toggl-cards/pkg/models"

type CreateDeckResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

func (dto *CreateDeckResponse) ModelToDTO(deck models.Deck) {
	*dto = CreateDeckResponse{
		DeckID:    deck.ID.String(),
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
}
